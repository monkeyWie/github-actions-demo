测试 github actions v1.0.13


[![Build Status](https://github.com/monkeyWie/github-actions-demo/workflows/build/badge.svg)](https://github.com/monkeyWie/github-actions-demo/actions?query=workflow%3Abuild)
[![Codecov](https://codecov.io/gh/monkeyWie/github-actions-demo/branch/master/graph/badge.svg)](https://codecov.io/gh/monkeyWie/github-actions-demo)

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
