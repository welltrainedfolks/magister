// Code generaTed by fileb0x at "2018-09-02 00:24:21.235038171 +0500 +05 m=+0.268671484" from config file "fileb0x.yml" DO NOT EDIT.
// modified(2018-06-02 12:05:13 +0500 +05)
// original path: assets/src/css/style.css

package assets

import (
  
  "os"
)

// FileStaticCSSStyleCSS is "static/css/style.css"
var FileStaticCSSStyleCSS = []byte("\x2e\x66\x6c\x61\x73\x68\x2d\x6d\x65\x73\x73\x61\x67\x65\x20\x7b\x0a\x20\x20\x20\x20\x62\x6f\x72\x64\x65\x72\x2d\x72\x61\x64\x69\x75\x73\x3a\x20\x35\x70\x78\x3b\x0a\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x3a\x20\x30\x2e\x34\x72\x65\x6d\x3b\x0a\x7d")

func init() {
  

  f, err := FS.OpenFile(CTX, "static/css/style.css", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    panic(err)
  }

  
  _, err = f.Write(FileStaticCSSStyleCSS)
  if err != nil {
    panic(err)
  }
  

  err = f.Close()
  if err != nil {
    panic(err)
  }
}

