测试 github actions v1.0.15


[![Build Status](https://github.com/monkeyWie/github-actions-demo/workflows/build/badge.svg)](https://github.com/monkeyWie/github-actions-demo/actions?query=workflow%3Abuild)
[![Codecov](https://codecov.io/gh/monkeyWie/github-actions-demo/branch/master/graph/badge.svg)](https://codecov.io/gh/monkeyWie/github-actions-demo)
[![Release](https://img.shields.io/github/release/monkeyWie/github-actions-demo.svg?style=flat-square)](https://github.com/monkeyWie/github-actions-demo/releases)

## Release

- create tag

```shell script
git tag $version
git push --tags
```

- delete tag
```shell script
git tag -d $version
git push --delete origin $version
```
