// Code generaTed by fileb0x at "2018-09-02 14:05:24.449655392 +0500 +05 m=+0.083407943" from config file "fileb0x.yml" DO NOT EDIT.
// modified(2018-09-02 14:03:58.570112749 +0500 +05)
// original path: assets/src/html/profile/skeleton.html

package assets

import (
  
  "os"
)

// FileProfileSkeletonHTML is "/profile/skeleton.html"
var FileProfileSkeletonHTML = []byte("\x3c\x73\x65\x63\x74\x69\x6f\x6e\x20\x63\x6c\x61\x73\x73\x3d\x22\x73\x65\x63\x74\x69\x6f\x6e\x22\x3e\x0a\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x63\x6f\x6c\x75\x6d\x6e\x73\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x63\x6f\x6c\x75\x6d\x6e\x20\x69\x73\x2d\x33\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x61\x73\x69\x64\x65\x20\x63\x6c\x61\x73\x73\x3d\x22\x6d\x65\x6e\x75\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x70\x20\x63\x6c\x61\x73\x73\x3d\x22\x6d\x65\x6e\x75\x2d\x6c\x61\x62\x65\x6c\x22\x3e\x47\x65\x6e\x65\x72\x61\x6c\x3c\x2f\x70\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x75\x6c\x20\x63\x6c\x61\x73\x73\x3d\x22\x6d\x65\x6e\x75\x2d\x6c\x69\x73\x74\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x6c\x69\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x61\x20\x63\x6c\x61\x73\x73\x3d\x22\x7b\x74\x61\x62\x2e\x67\x65\x6e\x65\x72\x61\x6c\x2e\x61\x63\x74\x69\x76\x65\x7d\x22\x20\x68\x72\x65\x66\x3d\x22\x2f\x70\x72\x6f\x66\x69\x6c\x65\x2f\x67\x65\x6e\x65\x72\x61\x6c\x2f\x22\x3e\x50\x72\x6f\x66\x69\x6c\x65\x20\x69\x6e\x66\x6f\x72\x6d\x61\x74\x69\x6f\x6e\x3c\x2f\x61\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x61\x20\x63\x6c\x61\x73\x73\x3d\x22\x7b\x74\x61\x62\x2e\x70\x61\x73\x73\x77\x6f\x72\x64\x2e\x61\x63\x74\x69\x76\x65\x7d\x22\x20\x68\x72\x65\x66\x3d\x22\x2f\x70\x72\x6f\x66\x69\x6c\x65\x2f\x70\x61\x73\x73\x77\x6f\x72\x64\x2f\x22\x3e\x50\x61\x73\x73\x77\x6f\x72\x64\x3c\x2f\x61\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x6c\x69\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x75\x6c\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x61\x73\x69\x64\x65\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x63\x6f\x6c\x75\x6d\x6e\x22\x20\x69\x64\x3d\x22\x70\x72\x6f\x66\x69\x6c\x65\x2d\x64\x61\x74\x61\x2d\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x22\x3e\x7b\x74\x61\x62\x2e\x64\x61\x74\x61\x7d\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x3c\x2f\x73\x65\x63\x74\x69\x6f\x6e\x3e")

func init() {
  

  f, err := FS.OpenFile(CTX, "/profile/skeleton.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    panic(err)
  }

  
  _, err = f.Write(FileProfileSkeletonHTML)
  if err != nil {
    panic(err)
  }
  

  err = f.Close()
  if err != nil {
    panic(err)
  }
}

