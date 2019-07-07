# Go Multi Modules Sample

## Usage

```
$ go build
$ ./go-multi-modules
```

## Directory structure

```
. // parent module
├── README.md // this file
├── go.mod // use `replace` directive
├── go.sum
├── main.go // import `gohome` and `helloworld` package
└── pkg
    ├── go-home // child module
    │   ├── go.mod
    │   ├── go.sum
    │   └── home.go
    └── hello-world
        └── hello-world.go
```