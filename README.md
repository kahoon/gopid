# ksql

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/kahoon/ksql/master/LICENSE)

A simple go utility to write a service pid to /var/run on linux systems

## Install

```
  go get github.com/kahoon/gopid
```
  
## Usage

```
package main

import (
       "github.com/kahoon/gopid"
)

func main() {
     err := gopid.Run("svcname") // will result is a file /var/run/svcname.pid to be created with the pid.
}

```
