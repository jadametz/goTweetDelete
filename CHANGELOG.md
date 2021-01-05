# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2021-01-05

### Added

- Ability to ignore deletion of a Tweet based on substring. See `IGNORESUBSTRINGS` configuration.

## [1.0.0] - 2020-12-06

### Changed

- **Breaking:** Support for ignoring multiple Tweet IDs via removal of `IGNOREID` in favor of `IGNOREIDS`

## [0.2.1] - 2020-11-13

### Fixed

- Fixed a bug ([#4](https://github.com/jadametz/goTweetDelete/issues/4)) where only the latest 200 Tweets were being retrieved

## [0.2.0] - 2020-09-12

### Added

- Ability to ignore deleting a specific Tweet

[1.1.0]: https://github.com/jadametz/gotweetdelete/releases/tag/v1.1.0
[1.0.0]: https://github.com/jadametz/gotweetdelete/releases/tag/v1.0.0
[0.2.1]: https://github.com/jadametz/gotweetdelete/releases/tag/v0.2.1
[0.2.0]: https://github.com/jadametz/gotweetdelete/releases/tag/v0.2.0
