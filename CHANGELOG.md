# Changelog

### v0.5.3
  * remove `omitempty` from the following:
    * `Collection.Quota` 
    * `Collection.ReadyForContent`
    * `FMD.Size`
    * `FMD.FormatValid`
    * `FMD.FormatAcceptable`
  * refactored tests for independence

### v0.5.2
  * correct `rsbe.Version` value
  * correct `JSON` `URL` tags in `IEListEntry` and `IEToSEEntry`
  * clean up code per `go.staticcheck` recommendations
  * add `LICENSE.txt`

### v0.5.1
  * add `RSBE API v0.7.0` development database dump needed for testing

### v0.5.0
  * add `Owner` functionality to align with `RSBE API v0.7.0`

### v0.4.0
  * add Golang module files

### v0.3.2
  * renamed RSBE database snapshot and moved to rsbe/testsupport
  * added config file that can be used to access an RSBE instance
    running on localhost

### v0.3.1
  * added RSBE database dump to `rsbe/testdata` to allow developers to
    set up an RSBE instance with the expected test data

### v0.3.0
  * added `rsbe.Version` constant
  * added Entity-to-Foreign ID (`EToFID`) functionality
  * requires [`rsbe`](https://github.com/nyudlts/rsbe) API `v0.1.1` or higher
