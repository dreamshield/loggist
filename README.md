Loggist
=======

Loggist is a simple and configurable log package.
It is natively written in the [Go](http://golang.org/) programming language. 

Features
------------------

* Xml configuring to be able to change logger parameters without recompilation
* Adjustable message formatting
* Different output writers
  * Console writer
  * File writer 
* Functions for flexible usage in libraries

Quick-start
-----------

```go
package main

import "github.com/dreamshield/loggist"

func main() {
	logger := loggist.NewLoggist(loggist.MODE_FILE, loggist.RECORD_MODE_YMDH)
	logger.Error("Error test")
	logger.Debugf("Debug test num=%d string=%s", 123, "abc")
}
```

Installation
------------

If you don't have the Go development environment installed, visit the 
[Getting Started](http://golang.org/doc/install.html) document and follow the instructions. Once you're ready, execute the following command:

The loggist has a default log configure file, to use customed configure you shuld do as bellow:

- make the LoggistConf directory
- modify the conf.xml file and copy it to the LoggistConf

```
go get -u github.com/dreamshield/loggist
```

Issues
---------------

Feel free to push issues that could make Loggist better:https://github.com/dreamshield/loggist/issues 

Changelog
---------------
* **v1.0** : Initial release. Features:
    * Xml config
    * Contraints and exceptions
    * Formatting
    * Receivers: console, file
