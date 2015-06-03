#goblog

**goblog** is a fast and simple static blog framework implemented by golang

##Features

- Amazing generating speed. Generating 1000 files(4KB per file) only need 5 seconds.
- Support GFM markdown (for more detail reference [mark](https://github.com/superhx/mark))
- Provide two ways **Command line** and **Web Page** to manage you blog

##Installation

``` bash
$ go install github.com/superhx/goblog
```

##Quick Start

###Set up blog workspace
``` bash
$ cd $blog_dir
$ goblog init
```

###New a blog
``` bash
$ goblog new "My first blog"
```

###Generate static files
``` bash
$ goblog generate
```

###Start the server
```bash
$ goblog server
```
