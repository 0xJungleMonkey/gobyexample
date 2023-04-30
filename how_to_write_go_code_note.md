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
~$ meow
Meow, Meow, Meow
```

## Go build

Go build only generate the executable in the same folder. And we need to run it with "./file".

```bash
~$ go build
~$ meow
zsh: command not found: meow

~$ ./meow
meow, meow, meow!
```

