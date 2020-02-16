# creep

A highly-specialized image download utility, useful for grabbing massive amounts of random images.

- [creep](#creep)
  - [Install](#install)
    - [Binaries](#binaries)
    - [Go Module](#go-module)
    - [Source](#source)
  - [Usage](#usage)
    - [Options](#options)
    - [Examples](#examples)
    - [Sample URLs](#sample-urls)
  - [Why](#why)
  - [Author](#author)
  - [License](#license)

## Install

### Binaries

### Go Module

### Source

## Usage

```
Usage:
  creep [FLAGS] [OPTIONS]

Options:
  -u, --url string        The URL of the resource to access (required)
  -c, --count int         The number of times to access the resource (defaults to 1)
  -n, --name string       The base filename to use as output (defaults to "creep")
  -o, --out string        The output directory path (defaults to current directory)
  -t, --throttle int      Number of seconds to wait between downloads (defaults to 0)

Flags:
  -h, --help              Prints help information
  -v, --version           Prints version information
```

### Options

`--url`

Specifies the HTTP URL of the image resource to access. This is the only required argument.

`--count`

The number of times to access and download a resource. Defaults to 1.

`--name`

The base filename of the downloaded resource. For example, given a `count` of `3`, a `name` of `cat` and `url` that returns `jpg`, `creep` will generate the follwing list of files:

```
cat-1.jpg
cat-2.jpg
cat-3.jpg
```

Defaults to "creep".

`--out`

The directory to save the output. If no directory is given, the current directory will be used. If the given directory does not exist, it will be created.

`--throttle`

Throttles downloads by the given number of seconds. Some URLs will return a given image based on the current time, so performing requests in very quick succession will yield duplicate images. If you're receiving duplicate images, it may help to throttle the download rate. Throttling is disabled by default.

### Examples

```bash
creep -u 
```

### Sample URLs

The following URLs will serve a random image upon request:

- Unsplash [https://source.unsplash.com/random](https://source.unsplash.com/random)
- This Person Does Not Exist [https://thispersondoesnotexist.com/image](https://thispersondoesnotexist.com/image)
- Picsum [https://picsum.photos/400](https://picsum.photos/400)
- Lorem Pixel [http://lorempixel.com/400](http://lorempixel.com/400)
- This Cat Does Not Exist [https://thiscatdoesnotexist.com/](https://thiscatdoesnotexist.com/)

## Why

## Author

[Christopher Murphy](https://github.com/Splode)

## License
