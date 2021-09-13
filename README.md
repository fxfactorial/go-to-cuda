This is a starter project repo to do golang with C++/Cuda

invoke make

nvcc -o libtest.so --shared -Xcompiler -fPIC test.cu
