// Code generaTed by fileb0x at "2018-09-02 00:24:21.2331079 +0500 +05 m=+0.266741150" from config file "fileb0x.yml" DO NOT EDIT.
// modified(2018-06-02 12:06:22 +0500 +05)
// original path: assets/src/html/partials/error.html

package assets

import (
  
  "os"
)

// FilePartialsErrorHTML is "/partials/error.html"
var FilePartialsErrorHTML = []byte("\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x68\x61\x73\x2d\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x2d\x77\x61\x72\x6e\x69\x6e\x67\x20\x66\x6c\x61\x73\x68\x2d\x6d\x65\x73\x73\x61\x67\x65\x22\x3e\x0a\x20\x20\x20\x20\x3c\x70\x3e\x54\x68\x65\x73\x65\x20\x65\x72\x72\x6f\x72\x73\x20\x77\x61\x73\x20\x65\x6e\x63\x6f\x75\x6e\x74\x65\x72\x65\x64\x3a\x3c\x2f\x70\x3e\x0a\x20\x20\x20\x20\x7b\x65\x72\x72\x6f\x72\x73\x7d\x0a\x20\x20\x20\x20\x3c\x70\x3e\x50\x6c\x65\x61\x73\x65\x20\x66\x69\x78\x20\x74\x68\x65\x6d\x20\x61\x6e\x64\x20\x74\x72\x79\x20\x61\x67\x61\x69\x6e\x2e\x3c\x2f\x70\x3e\x0a\x3c\x2f\x64\x69\x76\x3e\x0a\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x69\x73\x2d\x63\x6c\x65\x61\x72\x66\x69\x78\x22\x3e\x3c\x2f\x64\x69\x76\x3e")

func init() {
  

  f, err := FS.OpenFile(CTX, "/partials/error.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    panic(err)
  }

  
  _, err = f.Write(FilePartialsErrorHTML)
  if err != nil {
    panic(err)
  }
  

  err = f.Close()
  if err != nil {
    panic(err)
  }
}

