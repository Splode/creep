# creep

A specialized image download utility, useful for grabbing massive amounts of random images.

<img src="./.github/banner_v1.1.png">

![Go](https://github.com/Splode/creep/workflows/Go/badge.svg?branch=master)

Creep can be used to generated gobs of random image data quickly given a single URL. It has no dependencies or requirements and is cross-platform.

- [creep](#creep)
  - [Install](#install)
    - [Prebuilt Binaries](#prebuilt-binaries)
    - [Build from Source](#build-from-source)
  - [Usage](#usage)
    - [Options](#options)
    - [Examples](#examples)
    - [Sample URLs](#sample-urls)
  - [Why](#why)
  - [Contributing](#contributing)
  - [Author](#author)
  - [License](#license)

## Install

### Prebuilt Binaries

Install a prebuilt binary from the [releases page](https://github.com/Splode/creep/releases/latest).

### Build from Source

```bash
go get github.com/splode/creep/cmd/creep
```

## Usage

Simply pass in a URL that returns an image to `creep` to download. Pass in a number of images, and `creep` will download them all concurrently.

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

Download `32` random images to the current directory.

```bash
 creep -u https://thispersondoesnotexist.com/image -c 32
```

Download `64` random images using the base filename `random` to the `downloads` folder, throttling the download rate to `3` seconds.

```bash
creep --url=https://source.unsplash.com/random --name=random --out=downloads --count=64 --throttle=3
```

Download a single random image to the current directory.

```bash
creep -u https://source.unsplash.com/random
```

### Sample URLs

The following URLs will serve a random image upon request:

- Unsplash [https://source.unsplash.com/random](https://source.unsplash.com/random)
- This Person Does Not Exist [https://thispersondoesnotexist.com/image](https://thispersondoesnotexist.com/image)
- Picsum [https://picsum.photos/400](https://picsum.photos/400)
- Lorem Pixel [http://lorempixel.com/400](http://lorempixel.com/400)
- This Cat Does Not Exist [https://thiscatdoesnotexist.com/](https://thiscatdoesnotexist.com/)

## Why

I frequently find myself needing to seed application data sets with lots of images for testing or demos. Given a few minutes searching for a tool, I wasn't able to find something that suited my requirements, so I built one.

Why Go and not simply script `curl` or python? Go's concurrency model makes multiple HTTP requests _fast_, and being able to compile to a single, cross-platform binary is handy. Besides, Go's cool.

## Contributing

Contributions are welcome! See [CONTRIBUTING](https://github.com/Splode/creep/blob/master/.github/CONTRIBUTING.md) for details.

## Author

[Christopher Murphy](https://github.com/Splode)

## License

[MIT](https://github.com/Splode/creep/blob/master/LICENSE)
