# RTME-BROWSER

Display informations from RTME table

based on: https://github.com/sapk/go-genesys

# Screenshots
<img src="https://user-images.githubusercontent.com/4052400/59727736-2bf76180-9237-11e9-9cde-a0607ddc8bdc.PNG" width="480">
<img src="https://user-images.githubusercontent.com/4052400/59728667-53e8c400-923b-11e9-8216-8ac4e52bed04.PNG" width="1024">



## Build
```
go mod vendor
(cd assets/ui/ && yarn build)
go generate ./...
gox -os="darwin" -ldflags "-s -w" -output="build/{{.Dir}}-{{.OS}}-{{.Arch}}"
gox -os="linux" -ldflags "-s -w" -output="build/{{.Dir}}-{{.OS}}-{{.Arch}}"
gox -os="windows" -ldflags "-s -w -H windowsgui" -output="build/{{.Dir}}-{{.OS}}-{{.Arch}}"
upx -7 ./build/rtme-browser-{windows,linux}-*
```

Windows specific: `go build -ldflags "-s -w -H windowsgui" -o build/rtme-browser-windows-amd64.exe` 
```
export V=$(git describe --abbrev=0 --tags)
rcedit "./build/rtme-browser-windows-amd64.exe" --set-icon "assets/ui/public/favicon.ico" --set-file-version "${V#?}"
```
## Sources

 - icon: https://www.iconfinder.com/icons/2044249/checklist_list_icon 
