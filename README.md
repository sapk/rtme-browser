# RTME-BROWSER

Display information from RTME table

## Build
```
go mod vendor
(cd assets/ui/ && yarn build)
go generate
gox -ldflags "-s -w" -output="build/{{.Dir}}-{{.OS}}-{{.Arch}}"
```

## Sources

 - icon: https://www.iconfinder.com/icons/2044249/checklist_list_icon 