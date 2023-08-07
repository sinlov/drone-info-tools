# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.31.0](https://github.com/sinlov/drone-info-tools/compare/1.30.0...v1.31.0) (2023-08-07)

### ✨ Features

* change DRONE_COMMIT_BRANCH and DRONE_BUILD_TRIGGER default value ([f22c8fa3](https://github.com/sinlov/drone-info-tools/commit/f22c8fa3705b2776d02e0323043d00fd738bbd94))

## [1.30.0](https://github.com/sinlov/drone-info-tools/compare/1.29.0...v1.30.0) (2023-08-04)

### ✨ Features

* let drone_info mock by MockDroneInfoByRefsAndNumber can use env for test case ([f165a235](https://github.com/sinlov/drone-info-tools/commit/f165a235aef5c124a300f03bb9aa4fb0cffcefdf))

## [1.29.0](https://github.com/sinlov/drone-info-tools/compare/1.28.0...v1.29.0) (2023-08-04)

### ✨ Features

* add drone_info.NameCliStepsDebug and Drone.Build.BuildDebug for support build ([beaabb2d](https://github.com/sinlov/drone-info-tools/commit/beaabb2d579bb930ed7cfa0054442f09b29a635e))

## [1.28.0](https://github.com/sinlov/drone-info-tools/compare/1.27.0...v1.28.0) (2023-08-03)

### ✨ Features

* add drone_info.EnvDroneRepoBranch and Dron.Build.RepoBranch support ([b22948a8](https://github.com/sinlov/drone-info-tools/commit/b22948a84fa7ec21769be3966093b8c2f342885a))

## [1.27.0](https://github.com/sinlov/drone-info-tools/compare/1.26.0...v1.27.0) (2023-08-01)

### ✨ Features

* add drone_info.MockDroneInfoByRefsAndNumber to mock more info ([f8ac8472](https://github.com/sinlov/drone-info-tools/commit/f8ac84720e5ca5436761e349d96048f184f5db6b))

## [1.26.0](https://github.com/sinlov/drone-info-tools/compare/1.25.0...v1.26.0) (2023-08-01)

### 🐛 Bug Fixes

* fix time format Location not setting error ([100bef18](https://github.com/sinlov/drone-info-tools/commit/100bef1864ce3e02f0588aa160102cbd074cfed5))

### ✨ Features

* add drone_info.MockDroneInfoDroneSystemRefs can check args ([984a14a4](https://github.com/sinlov/drone-info-tools/commit/984a14a47d973ccec0d9e709ea0b3026b19667f7))

## [1.25.0](https://github.com/sinlov/drone-info-tools/compare/v1.24.0...v1.25.0) (2023-07-31)

### ✨ Features

* change drone_log ShowLogLineNo use runtime.Caller to get code lineno ([183d83d](https://github.com/sinlov/drone-info-tools/commit/183d83db8bde6b9c97430050112db3ba65d44135))

## [1.24.0](https://github.com/sinlov/drone-info-tools/compare/v1.23.0...v1.24.0) (2023-07-31)

### ✨ Features

* add drone_log.ShowLogLineNo and let OpenDebug open log shortfile flag ([bf284e2](https://github.com/sinlov/drone-info-tools/commit/bf284e282f6386f81bd8f900f5bb3094bb707386))

* add Drone.Build.SourceBranch and some env doc ([952e19c](https://github.com/sinlov/drone-info-tools/commit/952e19cb0e08d54b0af21bbc20353a1f4a6c983e))

## [1.23.0](https://github.com/sinlov/drone-info-tools/compare/v1.22.0...v1.23.0) (2023-07-18)

### ✨ Features

* add drone_log.Error and drone_log.Errorf fast print ([6d5ecd8](https://github.com/sinlov/drone-info-tools/commit/6d5ecd83ab02fceb9f1282cd0cd5669703387fd3))

## [1.22.0](https://github.com/sinlov/drone-info-tools/compare/v1.21.0...v1.22.0) (2023-07-18)

### ✨ Features

* add exit_cli for urfave_cli and drone_log for drone ([a61a57a](https://github.com/sinlov/drone-info-tools/commit/a61a57aa0cead7aaa37f6eb2c78600afb5f7321a))

## [1.21.0](https://github.com/sinlov/drone-info-tools/compare/v1.20.0...v1.21.0) (2023-07-05)

### ✨ Features

* add GetPackageJsonVersionGoStyleKeepSelect for pkgJson ([08777df](https://github.com/sinlov/drone-info-tools/commit/08777df5d619d6c823188d8ca3d21c4d3849c90e))

## [1.20.0](https://github.com/sinlov/drone-info-tools/compare/v1.10.0...v1.20.0) (2023-07-05)

### ✨ Features

* change to go1.18 to fix build ([b25a58e](https://github.com/sinlov/drone-info-tools/commit/b25a58e51c1836380faeb2dac5a01b093bfdc0ca))

* update safe depend and update code for safe ([fb99bea](https://github.com/sinlov/drone-info-tools/commit/fb99bea4223d5751f57da19e15ebfc3282962c02))

## [1.10.0](https://github.com/sinlov/drone-info-tools/compare/v1.9.0...v1.10.0) (2023-06-14)

### Features

* add env key drone_info.EnvKeyPluginDebug ([cffa4bc](https://github.com/sinlov/drone-info-tools/commit/cffa4bc11439a71972e6489c13d9f493e85720b7))

## [1.9.0](https://github.com/sinlov/drone-info-tools/compare/v1.8.0...v1.9.0) (2023-06-13)

### Features

* add drone_info.MockDroneInfoDroneSystemRefs can mock drone system info ([eaa10e2](https://github.com/sinlov/drone-info-tools/commit/eaa10e23cf4cfb963dbe9fc6c9e353ad8e5317a2))

## [1.8.0](https://github.com/sinlov/drone-info-tools/compare/v1.7.0...v1.8.0) (2023-02-11)

### Features

* add Repo.Link and get from Dron and mock test case ([c70a293](https://github.com/sinlov/drone-info-tools/commit/c70a2932bf6aa82aac3675f189a5042c3adff13b))

* mark v1.8.0 ([69bf5a7](https://github.com/sinlov/drone-info-tools/commit/69bf5a7c15e079d854c33b94ace187fd91de2e28))

## [1.7.0](https://github.com/sinlov/drone-info-tools/compare/v1.6.0...v1.7.0) (2023-02-11)

### Features

* add Drone.Build.Branch and function MockDroneInfoRefs to mock tags ([b9fc06a](https://github.com/sinlov/drone-info-tools/commit/b9fc06a7ba484e555af544f4d21ccae3b832a7d5))

## [1.6.0](https://github.com/sinlov/drone-info-tools/compare/v1.5.0...v1.6.0) (2023-02-08)

### Features

* change to z-MakefileUtils to reduce the impact on the project directory ([5697924](https://github.com/sinlov/drone-info-tools/commit/5697924b4a8c7b2f7598ded5baa2f1313700128b))

## [1.5.0](https://github.com/sinlov/drone-info-tools/compare/v1.4.1...v1.5.0) (2023-02-04)

### Features

* add Path2WebPath to let local path to web path ([8faaf06](https://github.com/sinlov/drone-info-tools/commit/8faaf06a7e76655a76fda51adfe05e5d373fbf98))

* make dist verison to v1.5.0 ([a168d99](https://github.com/sinlov/drone-info-tools/commit/a168d996dd40e54628c2e59f03b874df90542541))

### [1.4.1](https://github.com/sinlov/drone-info-tools/compare/v1.4.0...v1.4.1) (2023-02-03)

## [1.4.0](https://github.com/sinlov/drone-info-tools/compare/v1.3.0...v1.4.0) (2023-02-03)

### Features

* add pkgJson tools to embed package.json file ([ea2027d](https://github.com/sinlov/drone-info-tools/commit/ea2027d83bae53f6d53d694544a0a18a1e4cf8b3))

## [1.3.0](https://github.com/sinlov/drone-info-tools/compare/v1.2.0...v1.3.0) (2023-02-03)

### Features

* add drone_info.DroneBuildStatusStatusOptSupport() and template.RenderStatusSupport() ([1af8c6a](https://github.com/sinlov/drone-info-tools/commit/1af8c6ae49d4c7e3a2e97c6a286ec0325355f864))

## [1.2.0](https://github.com/sinlov/drone-info-tools/compare/v1.1.0...v1.2.0) (2023-02-03)

### Features

* add github.com/sinlov/drone-info-tools/tools for common use ([babb07c](https://github.com/sinlov/drone-info-tools/commit/babb07cbaf701da1921f324873863010baf89e05))

## [1.1.0](https://github.com/sinlov/drone-info-tools/compare/v1.0.1...v1.1.0) (2023-02-03)

### Features

* add module template and test case to use github.com/aymerick/raymond ([1b163c0](https://github.com/sinlov/drone-info-tools/commit/1b163c047f6e89c246ecadbd9756cafb4cb177b4))

## 1.0.0 (2023-02-02)

* first commit v1.0.0 ([2846922](https://github.com/sinlov/drone-info-tools/commit/2846922))
