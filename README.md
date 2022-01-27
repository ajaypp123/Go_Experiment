# Go_Experiment
Go Experiment and tutorial

# Install Go Package

- Install Package
```
rm -rf /usr/local/go
wget https://go.dev/dl/go1.17.6.linux-amd64.tar.gz .
tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz
```

- Set Path `/etc/environment`
```
PATH=$PATH:/usr/local/go/bin
```

- Test
```
 go version
```

- vscode setting
```
"gopls": {
    "experimentalWorkspaceModule": true,
}
```

# Create Go Development Env

- Update in `bashrc` or `/etc/environment`
  - `GOPATH=/home/ajayp/golib` will install all in build package.
  - `GOPATH=/home/ajayp/GoProject` is used for developer workspace
```
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin

export GOPATH=/home/ajayp/golib
export PATH=$PATH:$GOPATH/bin
export GOPATH=/home/ajayp/GoProject
```

- Create Project structure

```
GoProject
  - pkg
  - bin
  - src
```

- clone git repository inside src
  - `go mod init` set application.
```
cd src
git clone github.com/user/app
cd app
go mod init github.com/user/app
```

- Go command
```
go run main.go
go build
go clean -i importpath
go run
go install // create compiled binary inside pkg
go mod download github.com/ajaypp123/goutils
```
