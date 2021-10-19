hello-go:
	go build -o assets/binaries/standard-go-hello assets/code/helloworld.go

hello-tiny:
	tinygo build -o assets/binaries/tiny-go-hello assets/code/helloworld.go

hello-arduino: 
	tinygo build -o assets/binaries/tiny-arduino-go-hello --target=arduino --gc=none assets/code/helloworld.go 

flash-hello-arduino: 
	tinygo flash --target=arduino --gc=none assets/code/helloworld.go 

flash-hello-arduino-gc: 
	tinygo flash --target=arduino assets/code/helloworld.go 

flash-hello-arduino-gc-scheduler: 
	tinygo flash --target=arduino --scheduler=tasks assets/code/helloworld.go 

flash-example2: 
	tinygo flash --target=arduino --scheduler=tasks assets/code/example2/example2.go 

log-example2:
	cat /dev/ttyACM1

dashboard-linux:
	rm -rf assets/code/html
	mkdir assets/code/html
	cp "assets/code/dashboard/wasm_exec.js" "assets/code/html/wasm_exec.js"
	cp "assets/code/dashboard/wasm.js" "assets/code/html/wasm.js"
	cp "assets/code/dashboard/mqtt.js" "assets/code/html/mqtt.js"
	cp "assets/code/dashboard/index.html" "assets/code/html/index.html"
	mkdir assets/code/html/styles
	cp "assets/code/dashboard/styles/general.css" "assets/code/html/styles/general.css"
	mkdir "assets/code/html/assets"
	cp -r assets/code/assets assets/code/html
	tinygo build --target=wasm -o assets/code/html/wasm.wasm -opt=s --no-debug assets/code/dashboard/wasm.go
	go run "assets/code/api/main.go"

build-go-wasm:
	GOOS=js GOARCH=wasm go build -o assets/code/binaries/wasmgo.wasm assets/code/dashboard/wasm.go 

build-tiny-go-wasm:
	tinygo build -o assets/code/binaries/tiny-wasmgo.wasm --target=wasm assets/code/dashboard/wasm.go 