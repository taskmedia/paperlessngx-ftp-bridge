# Changelog

## [3.0.0](https://github.com/taskmedia/paperlessngx-ftp-bridge/compare/v2.1.1...v3.0.0) (2026-10-07)


### ⚠ BREAKING CHANGES

* multi-account-ready FTP config schema with PASV range and allowed extensions ([#91](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/91))
* embed FTP server, forward uploads directly to paperless-ngx ([#88](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/88))
* use cron to define interval ([#17](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/17))

### Features

* add Dockerfile ([4d81613](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/4d81613ef50dec88163c44a1ffdeb3e72152eabc))
* add ftpPath ([5a37da7](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/5a37da74323ae6d8ce2a79d5d955c264208943a7))
* add logLevel ([#14](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/14)) ([9f773ed](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/9f773edca336c7e0c1ca7a4d719ca63ca0a60b55))
* allow to use existing secret for ftp / pl password ([#42](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/42)) ([69e380c](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/69e380c69dcb8b09541519875ed19802668eeb1f))
* embed FTP server, forward uploads directly to paperless-ngx ([#88](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/88)) ([fc76bac](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/fc76bac87ce3878b984e803a31cf35946a051417))
* health check ([#16](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/16)) ([c9bd88e](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c9bd88e84d386d230a1e8f4b6e7ff536652eaf34))
* helm service topology, container hardening, probe wiring ([#92](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/92)) ([c6b59f3](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c6b59f30af69622cd0bdf244615cb9ea562463e3)), closes [#86](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/86)
* multi-account-ready FTP config schema with PASV range and allowed extensions ([#91](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/91)) ([88f3032](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/88f303266c3cce49f1924ad53655bc6838671d23)), closes [#84](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/84)
* offer explicit FTPS (AUTH TLS) with self-signed default and hot-reloadable existing cert ([#89](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/89)) ([7274802](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/72748022ea779d448a90eb0e91b959dac024e792)), closes [#83](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/83)
* re-enable ct install in Helm chart-testing CI ([#93](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/93)) ([ad1787c](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/ad1787ca50fb7e4ea4977301363964b92d0116b5)), closes [#87](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/87)
* release via release-please with draft-then-publish ([#107](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/107)) ([1c9e187](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/1c9e18766627b3d5d1d5f36b53159db375c0d9ab))
* skip latest/main-branch mutation for prerelease tags ([#95](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/95)) ([dec9edb](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/dec9edbcccb3b66fdbbc60003f49e0796b09132b))
* split health checks into /healthz liveness and /readyz readiness ([#90](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/90)) ([d6a7761](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/d6a7761a51bf795b998a30def68b2d78a060ae01)), closes [#85](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/85)
* use cron to define interval ([#17](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/17)) ([c563d90](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c563d9054e3b6e60239afdca810999faa0763508))
* use deployment with interval instead cronjob ([#10](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/10)) ([c8d1a94](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c8d1a94ba40796859e15c6d2cf52486b53c9bc1f))


### Bug Fixes

* **ci:** add needs to release wf ([#2](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/2)) ([fdae9f3](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/fdae9f3b97dd9210b6d0dab7dce0b9d116bf5fab))
* **ci:** fix package-name for delete pkg ([6e01ce3](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/6e01ce36f4961c70f8e9ad3cc3a8a2cd2301e815))
* cross-compile natively in Docker build instead of under QEMU ([#96](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/96)) ([f4a2a97](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/f4a2a9748c711d99b6477af84996dc4ed528e006))
* ensure pods are restarted on password update ([#21](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/21)) ([242e2ce](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/242e2ceed78ea6f88ee04b6049601e1d60f5d016))
* handle unchecked error returns for errcheck lint compliance ([67b856a](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/67b856afe6c0f8b2e6a1ef659fca3337e5aaad06))
* install golangci-lint via go install, not install.sh ([6498ebc](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/6498ebc92c78e8f06e707a28848417463a49d9f3))
* look up draft release id instead of relying on undocumented output ([683de61](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/683de61e395d401c8d3072b7ebafbf1fae18f056))
* package naming for image / chart ([#4](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/4)) ([3ed14bc](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/3ed14bc0b622fb0be32cbb8ff974a6e77cf178f8))
* remove usage of service account ([#7](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/7)) ([2700635](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/2700635b1669664d749c73bd5433c862840f768a))
* set chart version from release tag before packaging ([2a1f5ac](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/2a1f5ace49642114842371a37697078d24cbd553))
* stop moving release tag, update latest tag after builds ([#106](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/106)) ([ab0b7aa](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/ab0b7aaa98595881918ec041084bf6158c01f0d3))
* use correct paperless.url value ([#9](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/9)) ([c091053](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c09105324a8f62bb4c9201814758363c0bf8612f))
* use v prefix in image.tag version ([#8](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/8)) ([2922f59](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/2922f598caf54d47d2d9b1e9ede0a16f4b563946))


### Documentation

* add changelog for releases up to v2.0.1 ([#110](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/110)) ([838f132](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/838f132c8e79683f52f73a36a117eea211deb256))
* add migration notes for the embedded FTP server rewrite ([#94](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/94)) ([9f7ed9b](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/9f7ed9bc2d2252ee40f3a6a90127a5a1cf9c5db7))
* add moved notice to README ([0e0ca83](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/0e0ca8374f475ddab0ff29f477abebf7958c92ec))
* **adr:** record component breakdown for embedded FTP server (ADR-0001) ([bed36c3](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/bed36c3718d40f9897b2aa0d1cdfba5b4331c99f)), closes [#74](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/74)
* **adr:** record config schema for embedded FTP server (ADR-0002) ([5c6ebd8](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/5c6ebd8c11995e58b77d23a6a6534f05a1145452)), closes [#75](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/75)
* **adr:** record health/readiness model for embedded FTP server (ADR-0004) ([8e474ad](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/8e474ad31f833e631000964cdf80e50a828983f1)), closes [#77](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/77)
* **adr:** record Helm Service topology for embedded FTP server (ADR-0006) ([a0945f6](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/a0945f6df785494ed188b8d6d4f2a11ada977ea6)), closes [#79](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/79)
* **adr:** record testing strategy for embedded FTP server (ADR-0005) ([9410771](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/9410771bc45e9eadcb64db374c3b584098a532c9))
* **adr:** record TLS certificate provisioning for embedded FTP server (ADR-0003) ([636ee6e](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/636ee6ee0c71d603d49d1ff411a48616fceb2bcd)), closes [#76](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/76)
* **agents:** add agent instructions and issue tracker/domain docs ([6828ecb](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/6828ecb4e30686821446037f54693c399452e9c4))


### Continuous Integration

* count ci, chore and docs commits in releases ([#114](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/114)) ([a307d4d](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/a307d4da5ffe0e1b0505d6740311fea7876633f5))
* open helm chart bump pr after release ([#113](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/113)) ([c050b57](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c050b5795a5bf08592ca38ffc9c127538c1ebd19))
* skip tests on release-please prs ([#112](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/112)) ([dcc4b03](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/dcc4b03e3be9f73587ba54be4e7ce8fbf6195d79))
* update golangci-lint from v1.61.0 to v2.11.4 and Go from 1.23.1 to 1.25 ([24584b2](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/24584b20a2d68a7cd3d2a45e0b9c5e9c7e64e0c6))


### Miscellaneous Chores

* add .env to gitignore ([9d0e8a2](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/9d0e8a2d66305d4fa8ab0da731d7286af3f7d855))
* add build image wf ([dd17a7d](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/dd17a7d31dd335512e46b3cbd0492a829654072d))
* add chart and image metadata ([#1](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/1)) ([46129cb](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/46129cb14935acdd465cb2c2de8b5b1de7fa040e))
* add generated script ([0801f16](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/0801f16b378b43e75792b9fcf3f4e6e1e80f5af9))
* add main chart template ([c75ea7f](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c75ea7fb9b341d1159047662b05ff57e9255a86f))
* add Makefile ([e7d8fcc](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/e7d8fcc75b8b4a34649a0715062c180aa98f1365))
* add README ([af8dd7c](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/af8dd7cdec24883818bbc58f0a7273bcdd3485ab))
* add release tagname to pushed tags ([90c4f40](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/90c4f409c181ae3a4b215da5170246921e76e6a5))
* allow to specify imagePullPolicy ([#20](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/20)) ([59b3229](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/59b322930b1cbfdc5b86901080ed114d9fbffaaf))
* bump go.mod to 1.27.1, golangci-lint to v2.13.2 ([a270292](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/a270292d7c0fe8c7c8d02f801153414e1594b907))
* catch if file is not pdf ([4cf1346](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/4cf1346645076dd0ab66308422258dd3b3434d70))
* **ci:** add dependabot config ([#3](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/3)) ([20bd9df](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/20bd9df0e841d8123621d7261a08d2a013309e72))
* **ci:** add test-golang wf ([1b17f7f](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/1b17f7f25a01973723b85796637554940c7042f0))
* **ci:** add test-helm wf ([822c02e](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/822c02e8bc79ca11e3a77e0827e524162a9fa930))
* **ci:** disable ct install testing ([#12](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/12)) ([ec54ba1](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/ec54ba1a8e8afee44b4a063bb85556dcc6160a98))
* **ci:** enable push for build-image wf ([9478711](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/9478711296cfeb73966db0a90265d392ae9d80b5))
* **ci:** ensure steps prior version-bump use updated tag ([c14d0c3](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c14d0c3073317ff56aa39daa26a09e2d32595854))
* create release wf ([5ccf567](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/5ccf5673179b821d57edd53dd6853940d728d6cf))
* **deps:** bump actions/checkout from 4 to 7 ([#70](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/70)) ([8e945eb](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/8e945eb0f871574ebc6fdcab35717a2018583489))
* **deps:** bump azure/setup-helm from 4 to 5 ([#60](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/60)) ([e3aa80b](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/e3aa80ba37db1c92df9bd2322d44d87bb3419f37))
* **deps:** bump codecov/codecov-action from 4 to 5 ([#28](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/28)) ([1eae0a1](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/1eae0a123ddd140a35c92a87ec74b9d77c548413))
* **deps:** bump codecov/codecov-action from 5 to 7 ([#101](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/101)) ([c03bb0b](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c03bb0bf271a4ec157f98b3ee9d98fb59b1c9e00))
* **deps:** bump docker/build-push-action from 6.10.0 to 6.13.0 ([#34](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/34)) ([494c67d](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/494c67d5f9f42cc5a365ae0d0e65f97994e975f8))
* **deps:** bump docker/build-push-action from 6.13.0 to 6.15.0 ([#41](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/41)) ([a13c04f](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/a13c04f252e852e191b173a0c5f79cf4db56ed07))
* **deps:** bump docker/build-push-action from 6.15.0 to 7.1.0 ([#63](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/63)) ([2762604](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/276260486e1c31bd8633665266dbe75f2f023370))
* **deps:** bump docker/build-push-action from 6.9.0 to 6.10.0 ([#27](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/27)) ([8ee221c](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/8ee221cc843cdada151af88e321f426dca3b83f8))
* **deps:** bump docker/build-push-action from 7.1.0 to 7.4.0 ([#67](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/67)) ([e8ecdfe](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/e8ecdfe9669f490c7afba29b430dde26e9a01098))
* **deps:** bump docker/login-action from 3.3.0 to 4.6.0 ([#71](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/71)) ([d4ea7fe](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/d4ea7fe6580de60fffb28bab0bdc1846c61bc6e0))
* **deps:** bump docker/metadata-action from 5.5.1 to 5.6.1 ([#26](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/26)) ([cbbf9ad](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/cbbf9adba044dd1ba948d3aac857bbb07f606005))
* **deps:** bump docker/metadata-action from 5.6.1 to 5.7.0 ([#37](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/37)) ([98c94c4](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/98c94c4e546bba09ca800a46d83069226fc3817b))
* **deps:** bump docker/metadata-action from 5.7.0 to 6.0.0 ([#61](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/61)) ([e1b7aeb](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/e1b7aeb077913dc4144a0e0ad25c333f3f4e8a35))
* **deps:** bump docker/setup-buildx-action from 3.10.0 to 4.4.1 ([#69](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/69)) ([9cb8b6a](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/9cb8b6a62a4abc3ff0bc72aa8e3193f7e72f98c7))
* **deps:** bump docker/setup-buildx-action from 3.7.1 to 3.8.0 ([#31](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/31)) ([4f3179e](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/4f3179e5ca576ebe38c33b6126df12c5d8869a0e))
* **deps:** bump docker/setup-buildx-action from 3.8.0 to 3.10.0 ([#39](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/39)) ([bf85799](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/bf8579954cd6bd413c95ea42e9a846b75b32717d))
* **deps:** bump github.com/go-resty/resty/v2 from 2.15.3 to 2.16.2 ([#23](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/23)) ([706c5d7](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/706c5d7832425566943b9613e780d8975aba7961))
* **deps:** bump github.com/go-resty/resty/v2 from 2.16.2 to 2.16.5 ([#35](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/35)) ([4fd7d6d](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/4fd7d6d59f4e62258be492be6ced46dcad060e43))
* **deps:** bump github.com/jlaffaye/ftp from 0.2.0 to 0.2.1 ([#64](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/64)) ([34470d7](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/34470d7cffaffd1091c87c7a4182efac6f2c4cee))
* **deps:** bump github.com/jlaffaye/ftp from 0.2.1 to 0.2.4 ([#66](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/66)) ([08a88e0](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/08a88e01e8bfc38d90d1d4cb06c52e02e5f6a25e))
* **deps:** bump golang from 1.23-alpine to 1.24-alpine ([#36](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/36)) ([3998fd7](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/3998fd76e1dbd4aba34c817091fbb93fde1f0749))
* **deps:** bump golang from 1.25-alpine to 1.27-alpine ([#65](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/65)) ([a5dbbd9](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/a5dbbd9d76f4f9b676686334da8f789e6c633c74))
* **deps:** bump golang.org/x/net from 0.27.0 to 0.33.0 ([#30](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/30)) ([ec4cc56](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/ec4cc564d28e2c0534959466f7c6fb02ec9c2e30))
* **deps:** bump golang.org/x/net from 0.53.0 to 0.55.0 ([#72](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/72)) ([aab17c6](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/aab17c61c48d77f4ed78f2606efae7c63516e73e))
* **deps:** bump helm/chart-testing-action from 2.6.1 to 2.7.0 ([#32](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/32)) ([9b6f8e6](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/9b6f8e6cb505eb5c09f1a4c5d624c3551cfe3079))
* **deps:** bump helm/kind-action from 1.10.0 to 1.12.0 ([#33](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/33)) ([b80285f](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/b80285f061f6fe3eb9dfaa9c0bae2d46eef1c5df))
* **deps:** bump helm/kind-action from 1.12.0 to 1.14.0 ([#58](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/58)) ([40556af](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/40556aff29adf3407546b7a08bef1e242031f0d3))
* **deps:** bump helm/kind-action from 1.14.0 to 1.15.0 ([#68](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/68)) ([558dd6e](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/558dd6e182ca8c5f637859db75807b616cd0ed23))
* **deps:** bump JamesIves/github-pages-deploy-action ([#29](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/29)) ([b5f0837](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/b5f08375c26ee13025dc8a1bf7e001ae1097da8c))
* **deps:** bump JamesIves/github-pages-deploy-action ([#38](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/38)) ([c65d1a7](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c65d1a7ca8607e629be161b77d6cae392e1f7198))
* **deps:** bump JamesIves/github-pages-deploy-action ([#62](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/62)) ([fbd55e3](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/fbd55e3243ef7cbdb5bf9f103ad3dff8f06aaf30))
* **deps:** bump sigstore/cosign-installer from 3.7.0 to 3.8.1 ([#40](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/40)) ([d22f472](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/d22f472cf353bc14efc49378936c2d399f00bdd9))
* **deps:** bump softprops/action-gh-release from 1 to 2 ([#24](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/24)) ([8b9b4ae](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/8b9b4ae4215db4989025d7ce37fe8d410ac9407e))
* **deps:** bump softprops/action-gh-release from 2 to 3 ([#103](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/103)) ([e823c08](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/e823c08196b61c09d5d762653250e9a5733c6d63))
* **deps:** bump stefanzweifel/git-auto-commit-action from 4 to 5 ([#11](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/11)) ([1ac41de](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/1ac41de081a249c8adbc8ae52311fd2749c3ebdd))
* increase versions to keep to 200 ([#6](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/6)) ([0210111](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/021011116aac7f6d3d7f1ebade9249a3e8d6320d))
* init gomod ([e750048](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/e7500484df825f9aa11d1ab905020c5bd4eb7c3d))
* init Helm chart ([06aba21](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/06aba2119931e6029753ad844206414fc68aee38))
* link readme to chart ([36aef8d](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/36aef8d19d895181b4933db63c411489f73fe50e))
* **lint:** check return value buf.ReadFrom ([4a602ef](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/4a602ef5e8dad361fe47adb0cd6e0abf44e91422))
* **lint:** check return value conn.Quit ([74e9d05](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/74e9d053a8868aaecf33df121861b60e56c96852))
* **lint:** use attrib for logging when possible ([#15](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/15)) ([2f98f53](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/2f98f53ac7a99bd7a3d1ed68d3390a3a666e1983))
* **main:** release 2.1.0 ([#108](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/108)) ([1a9ec12](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/1a9ec129ebb2c0ca4b57bc448fee11eaf4b90e82))
* **main:** release 2.1.1 ([#115](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/115)) ([8083374](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/808337475e5644693f2bf3c76b482d871e6ff113))
* move helm chart to taskmedia/helm ([#105](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/105)) ([8bd3b75](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/8bd3b75fbd01d3fa704f61794cf78444a7851948))
* **release:** update version to v0.1.2 ([e16e5a8](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/e16e5a8608aae0e82259b9bc83bf7ed6bfcc59fe))
* **release:** update version to v0.1.3 ([efc5bda](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/efc5bda70c715db725f420d97bc442ecfea48760))
* **release:** update version to v0.1.4 ([5467983](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/5467983b196852db08ebddee7b0f423619af4902))
* **release:** update version to v0.1.5 ([0b1d3af](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/0b1d3af4ac221f87501c505ec17aeeb396c51ddb))
* **release:** update version to v0.1.6 ([d8f2cf5](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/d8f2cf544a017fe5172c166304e5b2f0832ad1ce))
* **release:** update version to v1.0.0 ([373feff](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/373feff8e0e530f5c36d153d7c0fbefd83951c4f))
* **release:** update version to v1.1.0 ([e202d37](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/e202d37b99b29bf77a5852e06e393a86547339dc))
* **release:** update version to v1.2.0 ([58b0b65](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/58b0b6583553cf47487551f602cc1a923b8b87bd))
* **release:** update version to v1.3.0 ([5ce1ca1](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/5ce1ca142e89e63aaf92f330f8cf7c2a9279eb69))
* **release:** update version to v1.3.1 ([a1f122e](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/a1f122e23e181ddc875c3afa5923ed2741fbb994))
* **release:** update version to v2.0.0 ([c438877](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c43887771a9bbdacd21d96a563c0c729b03c5fc4))
* rename base container ([#19](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/19)) ([4be545a](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/4be545a5bf7b431bc7f84cce718ccb0aa47f946c))
* rename chart directory ([1b1d6aa](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/1b1d6aa6e9bf416cdc0d96a3fe38d15ba2941099))
* rename envs to paperless ([b0630ad](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/b0630ad7f81e1a90f07ea9a11878cf361942e40b))
* rename ftpServer to ftpHost ([56cbc3e](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/56cbc3e770f28491922d50850d55bf5d5b71cd33))
* rename ftpUser to ftpUsername ([0620dee](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/0620deef5ad4912bf66dafe9be54b87c318701bc))
* rename module ([f676204](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/f67620436800eb0b50ce3a0daa63d4c86055d6a5))
* rename release wf ([297f3fd](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/297f3fd6073460cdeea591c384cdc6bc974ff9f6))
* rename repo without helm prefix ([#13](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/13)) ([049c3b5](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/049c3b58bce7023a86b9f6fe6c64b9b73c46075c))
* required chart fields ([#18](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/18)) ([c1a1018](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c1a1018ebe62917b47cba64ba37ce3b2bf6f4278))
* rm Helm test ([1f059b5](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/1f059b5d020c86982f69422e4dbf7722235775c7))
* tls explicit ([3cbe444](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/3cbe444167188cb5df55d2b9ae9e19c891e4cbd7))
* upload binaries on release ([#22](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/22)) ([875dff5](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/875dff501f8edabebd442dd73b9b5668d6dbfbb0))
* use release-please-action v5 ([#109](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/109)) ([2e6623f](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/2e6623fe4c09cf6a63c66baaec4308da19361324))
* use the chart version as default image tag ([#5](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/5)) ([64202af](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/64202afc766994d00877c81be2115c85466dbcbf))
* use tls for ftp ([c4245ae](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/c4245aeb9148661212bb9fd44dc3383dbee4feec))
* use username/password instead api token for pl auth ([7b220fa](https://github.com/taskmedia/paperlessngx-ftp-bridge/commit/7b220fa9581f7a44837171f8eee5b0b9c24bd2ba))

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
