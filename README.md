# netshare
![GitHub](https://img.shields.io/github/license/Leixb/netshare)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/Leixb/netshare)
[![Build Status](https://travis-ci.com/Leixb/netshare.svg?branch=master)](https://travis-ci.com/Leixb/netshare)
[![Go Report Card](https://goreportcard.com/badge/github.com/Leixb/netshare)](https://goreportcard.com/report/github.com/Leixb/netshare)
[![GoDoc](https://godoc.org/github.com/Leixb/netshare?status.svg)](https://godoc.org/github.com/Leixb/netshare)

Serve and recieve static files on local network. 

```
usage: ./netshare [-h|--help] [-u|--upload] [-f|--folder "<value>"] [-p|--port
                  <integer>] [-d|--debug]

                  Share folder on local network

Arguments:

  -h  --help    Print help information
  -u  --upload  Allow upload to folder
  -f  --folder  Folder to share. Default: .
  -p  --port    Port to use. Default: 8080
  -d  --debug   Gin in debug mode
```
