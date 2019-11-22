# godll
Example for create dll from golang


1. install TDM-GCC from http://tdm-gcc.tdragon.net/download
2. run go build -ldflags="-s -w" -buildmode=c-shared -o main.dll src\github.com\vchatchai\godll\main.go
3. the result  has 2 files 
    1. main.dll 
    2. main.h


Thank you
https://zhuanlan.zhihu.com/p/46721768