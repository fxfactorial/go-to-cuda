# Always keep in mind: `go doc cmd/cgo|less`

all:
	nvcc -o libtest.so --shared -Xcompiler -fPIC test.cu
	CGO_ENABLED=1 go build -o m
	GODEBUG=cgocheck=11 ./m

clean:; rm -rf m libtest.so
