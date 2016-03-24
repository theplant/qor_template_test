# App

```
$ glide i
$ go run main.go &
$ curl -s -i http://localhost:9999/ | head -n 1
2016/03/24 16:46:16 Finish [GET] / Took 4.41ms
HTTP/1.1 200 OK
```

# Test

```
$ go test -v ./app
=== RUN   TestApp
2016/03/24 16:46:48 Finish [GET] / Took 0.15ms
2016/03/24 16:46:48 http: panic serving 127.0.0.1:64013: runtime error: invalid memory address or nil pointer dereference
goroutine 18 [running]:
net/http.(*conn).serve.func1(0xc820298000)
	/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:1389 +0xc1
panic(0x54fa60, 0xc820012100)
	/usr/local/Cellar/go/1.6/libexec/src/runtime/panic.go:426 +0x4e9
html/template.(*Template).escape(0x0, 0x0, 0x0)
	/usr/local/Cellar/go/1.6/libexec/src/html/template/template.go:79 +0x3c
html/template.(*Template).Execute(0x0, 0xbc01e0, 0xc82029a0d0, 0x5ee9e0, 0xc820290070, 0x0, 0x0)
	/usr/local/Cellar/go/1.6/libexec/src/html/template/template.go:101 +0x34
github.com/theplant/qor_template_test/vendor/github.com/qor/admin.(*Context).Execute(0xc820290070, 0x608060, 0x9, 0x0, 0x0)
	$GOPATH/src/github.com/theplant/qor_template_test/vendor/github.com/qor/admin/context.go:194 +0xbab
github.com/theplant/qor_template_test/vendor/github.com/qor/admin.(*controller).Dashboard(0xc820226a90, 0xc820290070)
	$GOPATH/src/github.com/theplant/qor_template_test/vendor/github.com/qor/admin/controller.go:21 +0x42
github.com/theplant/qor_template_test/vendor/github.com/qor/admin.(*controller).Dashboard-fm(0xc820290070)
	$GOPATH/src/github.com/theplant/qor_template_test/vendor/github.com/qor/admin/route.go:106 +0x2a
github.com/theplant/qor_template_test/vendor/github.com/qor/admin.(*Admin).compile.func2(0xc820290070, 0xc8201f9640)
	$GOPATH/src/github.com/theplant/qor_template_test/vendor/github.com/qor/admin/route.go:324 +0x919
github.com/theplant/qor_template_test/vendor/github.com/qor/admin.Middleware.Next(0x607f30, 0xa, 0xc820226b50, 0xc8201f9640, 0xc820290070)
	$GOPATH/src/github.com/theplant/qor_template_test/vendor/github.com/qor/admin/route.go:35 +0x36
github.com/theplant/qor_template_test/vendor/github.com/qor/admin.(*Admin).compile.func1(0xc820290070, 0xc8201f9620)
	$GOPATH/src/github.com/theplant/qor_template_test/vendor/github.com/qor/admin/route.go:294 +0xbc
github.com/theplant/qor_template_test/vendor/github.com/qor/admin.(*Admin).ServeHTTP(0xc8201adf40, 0xbc0108, 0xc82029a0d0, 0xc8202b0000)
	$GOPATH/src/github.com/theplant/qor_template_test/vendor/github.com/qor/admin/route.go:370 +0x730
net/http.(*ServeMux).ServeHTTP(0xc8201b3050, 0xbc0108, 0xc82029a0d0, 0xc8202b0000)
	/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:1910 +0x17d
net/http.serverHandler.ServeHTTP(0xc820252300, 0xbc0108, 0xc82029a0d0, 0xc8202b0000)
	/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:2081 +0x19e
net/http.(*conn).serve(0xc820298000)
	/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:1472 +0xf2e
created by net/http.(*Server).Serve
	/usr/local/Cellar/go/1.6/libexec/src/net/http/server.go:2137 +0x44e
--- FAIL: TestApp (0.00s)
	app_test.go:14: Get http://127.0.0.1:64012: EOF
FAIL
exit status 1
FAIL	github.com/theplant/qor_template_test/app	0.016s
```

