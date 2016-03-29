package test

import "os"

var cwd string

func init() {
    cwd, _ = os.Getwd()
}
