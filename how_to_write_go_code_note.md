# Note

for [how to write go code](https://go.dev/doc/code) & [how to write go code with gopath](https://go.dev/doc/gopath_code)

## Gopath & path

Gopath is an env variable.
Path is an env variable as well. 

Once we define the Gopath & Path in ~/.zshrc, they are defined in the terminal. When we run the executable file anywhere, the terminal can recognize it and run the program.

## Go install

We need to run it in the module directory.  It will build and generate a file in ./go/bin. You can simply run the module anywhere.

```bash
~$ go install
(or
~$ go install github.com/0xJungleMonkey/meow)
~$ meow
Meow, Meow, Meow
```

## Go build

Go build only generate the executable in the same folder.
Go build can test the package compiles, will not compile if there is any error.
And we need to run it with "./file".

```bash
~$ go build
(or 
~$ go build github.com/0xJungleMonkey/meow)
~$ meow
zsh: command not found: meow

~$ ./meow
meow, meow, meow!
```


## Program vs Library

Program is executable, to finish certain task, not meant for reuse. We will run go install to make it a command executable. 

Library is polyphizmed, we wants to reuse it.


## GOROOT

If you use GOPATH, you need to turn off GOROOT:
```bash
go env -w GO111MODULE=off
```

