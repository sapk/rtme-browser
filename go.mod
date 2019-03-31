module github.com/sapk/rtme-browser

go 1.12

require (
	github.com/gin-contrib/gzip v0.0.1
	github.com/gin-gonic/gin v1.3.1-0.20190328061400-ce20f107f5dc
	github.com/mattn/go-colorable v0.1.1
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/rs/zerolog v1.13.0
	github.com/sapk/go-genesys v0.0.0-20190331030831-4ee8281b8536
)

replace github.com/golang/lint => golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3

replace github.com/go-xorm/xorm => github.com/techknowlogick/xorm v0.7.2-0.20190330194841-10f0de7dc384
