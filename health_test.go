package main

import (
	"errors"
	"net/http/httptest"
	"testing"
	"time"
)

type fakeFTPStatus struct {
	addr string
}

func (f *fakeFTPStatus) Addr() string { return f.addr }

type fakePaperlessPinger struct {
	err error
}

func (f *fakePaperlessPinger) Ping() error { return f.err }

func TestReadinessCheckerIsLive(t *testing.T) {
	tests := []struct {
		desc string
		addr string
		want bool
	}{
		{desc: "listener bound", addr: "127.0.0.1:2121", want: true},
		{desc: "listener not bound", addr: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			checker := newReadinessChecker(&fakeFTPStatus{addr: tt.addr}, &fakePaperlessPinger{}, time.Hour)

			if got := checker.isLive(); got != tt.want {
				t.Errorf("isLive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadinessCheckerIsReadyReflectsBothListenerAndCachedPaperlessCheck(t *testing.T) {
	tests := []struct {
		desc            string
		ftpAddr         string
		paperlessErr    error
		skipInitialPing bool
		want            bool
	}{
		{desc: "listener bound and paperless reachable", ftpAddr: "127.0.0.1:2121", paperlessErr: nil, want: true},
		{desc: "listener bound but paperless unreachable", ftpAddr: "127.0.0.1:2121", paperlessErr: errors.New("boom"), want: false},
		{desc: "listener not bound even though paperless reachable", ftpAddr: "", paperlessErr: nil, want: false},
		{desc: "no background check has run yet", ftpAddr: "127.0.0.1:2121", skipInitialPing: true, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			checker := newReadinessChecker(&fakeFTPStatus{addr: tt.ftpAddr}, &fakePaperlessPinger{err: tt.paperlessErr}, time.Hour)

			if !tt.skipInitialPing {
				checker.checkPaperless()
			}

			if got := checker.isReady(); got != tt.want {
				t.Errorf("isReady() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadinessCheckerRunIsNonFatalOnInitialPaperlessFailure(t *testing.T) {
	checker := newReadinessChecker(&fakeFTPStatus{addr: "127.0.0.1:2121"}, &fakePaperlessPinger{err: errors.New("down")}, time.Hour)

	done := make(chan struct{})
	go func() {
		checker.checkPaperless()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("checkPaperless() did not return")
	}

	if checker.isReady() {
		t.Errorf("isReady() = true, want false after an unreachable paperless-ngx check")
	}
	if !checker.isLive() {
		t.Errorf("isLive() = false, want true: an unreachable paperless-ngx must not affect liveness")
	}
}

func TestHealthzHandlerReflectsListenerState(t *testing.T) {
	tests := []struct {
		desc       string
		addr       string
		wantStatus int
	}{
		{desc: "listener bound", addr: "127.0.0.1:2121", wantStatus: 200},
		{desc: "listener not bound", addr: "", wantStatus: 503},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			checker := newReadinessChecker(&fakeFTPStatus{addr: tt.addr}, &fakePaperlessPinger{}, time.Hour)

			rec := httptest.NewRecorder()
			healthzHandler(checker)(rec, httptest.NewRequest("GET", "/healthz", nil))

			if rec.Code != tt.wantStatus {
				t.Errorf("healthzHandler() status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestReadyzHandlerReflectsListenerAndPaperlessState(t *testing.T) {
	tests := []struct {
		desc         string
		addr         string
		paperlessErr error
		wantStatus   int
	}{
		{desc: "both healthy", addr: "127.0.0.1:2121", paperlessErr: nil, wantStatus: 200},
		{desc: "paperless unreachable", addr: "127.0.0.1:2121", paperlessErr: errors.New("boom"), wantStatus: 503},
		{desc: "listener not bound", addr: "", paperlessErr: nil, wantStatus: 503},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			checker := newReadinessChecker(&fakeFTPStatus{addr: tt.addr}, &fakePaperlessPinger{err: tt.paperlessErr}, time.Hour)
			checker.checkPaperless()

			rec := httptest.NewRecorder()
			readyzHandler(checker)(rec, httptest.NewRequest("GET", "/readyz", nil))

			if rec.Code != tt.wantStatus {
				t.Errorf("readyzHandler() status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}
