# GolangReadingGroup

練習三種golang執行外部function的方法

## 同時compier
在basic_exam中的main.go有使用到addone function，此function在function.go中定義

```
cd basic_exam
go build
./basic_exam
```
## 使用package
將function複製到自己的GOROOT路徑下，可以使用指令go env知道自己的GOROOT 
例如：GOROOT="/usr/local/go" 
`sudo cp -R  mypkg /usr/local/go/src/` 
使用時就可import使用 

```
cd pack_exam
go build
./pack_exam
```
package內若要給外部使用的function，必須是**大寫開頭**


## 包裝成plugin
把function經過編譯之後，可以方便在保護原始碼的情況下，重複使用 
可以只重新build plugin但是main program不重複compiler 

```
cd plugin_exam
go build -buildmode=plugin plugin.go #產生plugin.so
go build
./plugin_exam
```