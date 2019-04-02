# RTME-BROWSER

Display informations from RTME table

based on: https://github.com/sapk/go-genesys

## Build
```
go mod vendor
(cd assets/ui/ && yarn build)
go generate
gox -os="darwin" -ldflags "-s -w" -output="build/{{.Dir}}-{{.OS}}-{{.Arch}}"
gox -os="linux" -ldflags "-s -w" -output="build/{{.Dir}}-{{.OS}}-{{.Arch}}"
gox -os="windows" -ldflags "-s -w -H windowsgui" -output="build/{{.Dir}}-{{.OS}}-{{.Arch}}"
upx -7 ./build/rtme-browser-{windows,linux}-*
```

Windows specific: `go build -ldflags "-s -w -H windowsgui" -o build/rtme-browser-windows-amd64.exe` 

## Sources

 - icon: https://www.iconfinder.com/icons/2044249/checklist_list_icon 
