# RTME-BROWSER

Display informations from RTME table

based on: https://github.com/sapk/go-genesys

## Build
```
go mod vendor
(cd assets/ui/ && yarn build)
go generate
gox -ldflags "-s -w" -output="build/{{.Dir}}-{{.OS}}-{{.Arch}}"
upx -7 ./build/rtme-browser-{windows,linux}-*
```

## Sources

 - icon: https://www.iconfinder.com/icons/2044249/checklist_list_icon 
