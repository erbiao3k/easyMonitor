 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" && upx easyMonitor && scp easyMonitor root@8.218.122.49:/
