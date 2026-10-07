# Changelog

## [2.1.1](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v2.1.0...v2.1.1) (2026-10-07)


### Continuous Integration

* count ci, chore and docs commits in releases ([#114](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/114)) ([a307d4d](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/a307d4da5ffe0e1b0505d6740311fea7876633f5))
* open helm chart bump pr after release ([#113](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/113)) ([c050b57](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c050b5795a5bf08592ca38ffc9c127538c1ebd19))
* skip tests on release-please prs ([#112](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/112)) ([dcc4b03](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/dcc4b03e3be9f73587ba54be4e7ce8fbf6195d79))

## [2.1.0](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v2.0.1...v2.1.0) (2026-10-07)


### Features

* release via release-please with draft-then-publish ([#107](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/107)) ([1c9e187](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/1c9e18766627b3d5d1d5f36b53159db375c0d9ab))


### Bug Fixes

* look up draft release id instead of relying on undocumented output ([683de61](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/683de61e395d401c8d3072b7ebafbf1fae18f056))


## [2.0.1](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v2.0.0...v2.0.1) (2026-10-07)


### Bug Fixes

* stop moving release tag, update latest tag after builds ([#106](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/106)) ([ab0b7aa](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/ab0b7aaa98595881918ec041084bf6158c01f0d3))


## [2.0.0](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v1.3.1...v2.0.0) (2026-10-07)


### ⚠ BREAKING CHANGES

* multi-account-ready FTP config schema with PASV range and allowed extensions
* embed FTP server, forward uploads directly to paperless-ngx


### Features

* skip latest/main-branch mutation for prerelease tags ([#95](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/95)) ([dec9edb](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/dec9edbcccb3b66fdbbc60003f49e0796b09132b))
* re-enable ct install in Helm chart-testing CI ([#93](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/93)) ([ad1787c](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/ad1787ca50fb7e4ea4977301363964b92d0116b5))
* helm service topology, container hardening, probe wiring ([#92](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/92)) ([c6b59f3](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c6b59f30af69622cd0bdf244615cb9ea562463e3))
* multi-account-ready FTP config schema with PASV range and allowed extensions ([#91](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/91)) ([88f3032](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/88f303266c3cce49f1924ad53655bc6838671d23))
* split health checks into /healthz liveness and /readyz readiness ([#90](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/90)) ([d6a7761](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/d6a7761a51bf795b998a30def68b2d78a060ae01))
* offer explicit FTPS (AUTH TLS) with self-signed default and hot-reloadable existing cert ([#89](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/89)) ([7274802](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/72748022ea779d448a90eb0e91b959dac024e792))
* embed FTP server, forward uploads directly to paperless-ngx ([#88](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/88)) ([fc76bac](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/fc76bac87ce3878b984e803a31cf35946a051417))


### Bug Fixes

* set chart version from release tag before packaging ([2a1f5ac](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/2a1f5ace49642114842371a37697078d24cbd553))
* install golangci-lint via go install, not install.sh ([6498ebc](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/6498ebc92c78e8f06e707a28848417463a49d9f3))
* cross-compile natively in Docker build instead of under QEMU ([#96](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/96)) ([f4a2a97](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/f4a2a9748c711d99b6477af84996dc4ed528e006))


## [1.3.1](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v1.3.0...v1.3.1) (2026-09-10)


## [1.3.0](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v1.2.0...v1.3.0) (2026-04-19)


### Bug Fixes

* handle unchecked error returns for errcheck lint compliance ([67b856a](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/67b856afe6c0f8b2e6a1ef659fca3337e5aaad06))


## [1.2.0](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v1.1.0...v1.2.0) (2025-03-13)


### Features

* allow to use existing secret for ftp / pl password ([#42](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/42)) ([69e380c](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/69e380c69dcb8b09541519875ed19802668eeb1f))


## [1.1.0](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v1.0.0...v1.1.0) (2024-11-04)


### ⚠ BREAKING CHANGES

* use cron to define interval


### Features

* use cron to define interval ([#17](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/17)) ([c563d90](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c563d9054e3b6e60239afdca810999faa0763508))
* health check ([#16](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/16)) ([c9bd88e](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c9bd88e84d386d230a1e8f4b6e7ff536652eaf34))
* add logLevel ([#14](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/14)) ([9f773ed](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/9f773edca336c7e0c1ca7a4d719ca63ca0a60b55))


### Bug Fixes

* ensure pods are restarted on password update ([#21](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/21)) ([242e2ce](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/242e2ceed78ea6f88ee04b6049601e1d60f5d016))


## [1.0.0](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v0.1.6...v1.0.0) (2024-10-30)


### Features

* use deployment with interval instead cronjob ([#10](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/10)) ([c8d1a94](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c8d1a94ba40796859e15c6d2cf52486b53c9bc1f))


## [0.1.6](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v0.1.5...v0.1.6) (2024-10-29)


### Bug Fixes

* use correct paperless.url value ([#9](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/9)) ([c091053](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c09105324a8f62bb4c9201814758363c0bf8612f))


## [0.1.5](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v0.1.4...v0.1.5) (2024-10-29)


### Bug Fixes

* use v prefix in image.tag version ([#8](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/8)) ([2922f59](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/2922f598caf54d47d2d9b1e9ede0a16f4b563946))
* remove usage of service account ([#7](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/7)) ([2700635](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/2700635b1669664d749c73bd5433c862840f768a))


## [0.1.4](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v0.1.3...v0.1.4) (2024-10-29)


### Bug Fixes

* package naming for image / chart ([#4](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/4)) ([3ed14bc](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/3ed14bc0b622fb0be32cbb8ff974a6e77cf178f8))


## 0.1.3 (2024-10-28)


### Features

* add ftpPath ([5a37da7](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/5a37da74323ae6d8ce2a79d5d955c264208943a7))
* add Dockerfile ([4d81613](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/4d81613ef50dec88163c44a1ffdeb3e72152eabc))


### Bug Fixes

* **ci:** add needs to release wf ([#2](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/2)) ([fdae9f3](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/fdae9f3b97dd9210b6d0dab7dce0b9d116bf5fab))
* **ci:** fix package-name for delete pkg ([6e01ce3](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/6e01ce36f4961c70f8e9ad3cc3a8a2cd2301e815))
