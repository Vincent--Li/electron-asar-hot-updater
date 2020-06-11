# Updater for electron-asar-hot-updater (Golang)

## What it is

> golang version of updater.cs which doesn't require .Net Framework on windowns 7

## How it work

1. Prerequisite.

> The golang environment. (Go 1.14.4 over my machine, others may work as well)

2. Run following command. You will get an updater.exe in current path

> go build -ldflags "-H windowsgui" -o updater.exe

-ldflags "-H windowsgui" : indicate that this program should run background instead of prompting a CMD window
-o : compile our program to specific file, here name is updater.exe

3. Ready to go

Everything is the same except for that you need to replace the "updater.exe" in root path of this project with the one you just compile.

## Notice
1. File nac.manifest / nac.syso only needed when your program require UAC on windows
2. You can generate nac.syso by way of 

```shell
go get github.com/akavel/rsrc
rsrc -manifest nac.manifest -o nac.syso
go build -ldflags "-H windowsgui" -o updater.exe
```

