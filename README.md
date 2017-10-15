# Todo

This is a command line interface app for todos. It is backed by a file. Each
line of the file represents a todo. Each line is in json format so it's
human readable even without the application. The file is stored in Unix systems
at `~/.todo`. This is currently not configurable (unless you build from source
and modify the `app.DataFilePath`). The output is intended be pipe-able to
other applications for further manipulation

<p align="center">
<img alt="usage screencast" src="https://user-images.githubusercontent.com/348407/31581016-b281e15e-b12d-11e7-8d95-50366f6af938.gif" />
</p>

## Installation

Pull down the repo and `go install`. Or pull down a release and put the binary
in a directory in your path.

## Usage

```
$  todo --help
Usage of todo:
  -a string
    	Add a todo using -a
  -c -c 0
    	Complete a todo by index -c 0. (default -1)
  -only-completed --only-completed
    	Show Only Completed Todos --only-completed
  -r -r 0
    	Remove a todo by index -r 0. (default -1)
  -show-completed --show-completed
    	Show Completed Todos --show-completed
```

## Contributing

Please open a pull request
