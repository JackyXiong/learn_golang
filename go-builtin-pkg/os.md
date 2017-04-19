### 概述
os 提供操作系统相关的接口，统一平台间的差异，不会提供部分操作系统特有的接口
* 主要是提供文件、目录和进程相关的操作
* 狠毒实现是对syscall package 的封装


### 类型
#### File
File type 对应的是系统的文件描述符，文件相关的操作，都是以程序的工作目录为默认目录
* os.Create 创建一个用于IO操作的File指针，可读可写。os.OpenFile 的封装，谨慎使用，需要先判断文件是否存在。
因为文件存在，使用该函数会将源文件截断为0长度，原有内容都删除。
* os.NewFile 由文件描述符返回File指针，不保存，并没有真正创建文件。
* os.Open 打开文件，只读，OpenFile的封装
* os.OpenFile 打开文件，通过参数可选不同的模式（只读、可读可写）和权限(mode)。函数实现是调用了syscall.Open和os.OpenFile
支持绝对路径和相对路径
* os.Pipe 返回可读的File指针r和可写的File指针w。写入到w，可从r中读取。
以下f是指一个`*File`
* f.Chdir 修改f的dir，参数需要是有效dir
* f.Chmod 修改f的mode，参数需要是有效mode(FileMode type)
* f.Close 关闭文件，之后无法进行io操作
* f.Fd 返回的是文件描述符
* f.Name 返回文件名
* f.Read 参数是b []byte(slice)，读取len(b)长度的f到b中
* f.ReadAt 参数是b 和 偏移量off，在偏移的基础上读取
* f.Write 参数是b []byte(slice)，读取b的内容写入到f。
* f.WriteAt 有偏移的基础上写入 
* f.WriteString(str) 写入字符串
NOTE:写入数据时会在start:end之间写入，其他的数据不会被影响。
```
f, _ := os.OpenFile("test", os.O_WRONLY|os.O_CREATE, os.ModePerm)
n, err := f.WriteString("1234")
f.Sync()
n, err = f.WriteString("ab")
```
最后test的内容是ab34
* f.Readdir(n int) 读取f所关联的目录下的文件信息，n>0读取前n个，n<=0读取所有。
返回的是FileInfo 的slice。f关联的需要是一个目录，函数才能正常使用，如果是文件，则返回的是nil
* f.Readdirnames(n)，返回的是string的slice，作用和Readdir 类似，但是只返回文件名
* f.Seek
* f.Stat 获取f的信息，返回FileInfo
* f.Sync 将内存中的文件内容保存到磁盘，即Save
* f.Truncate(n) 修改文件长度到n，其后内容舍弃


#### FileInfo 
文件信息类型，提供Name(), Size(), Mode(), ModTime(), IsDir(), Sys()
相关函数和方法
* os.Lstat(name) name是文件名，返回FileInfo。如果是软链接，直接返回软链接的信息
* os.Lstat(name) name是文件名，返回FileInfo。如果是软链接，返回被文件所链接的文件的信息

#### FileMode
实际上是unit32 type，表示文件模式和权限。在os包内，被用作定义常量
相关函数和方法
* IsDir()
* IsRegular()
* Perm()
* String()

#### PathError
由路径引起的错误。封装error的struct，提供了Error方法

#### LinkError
由硬链接，软连接和重命名系统调用引起的错误。封装error的strcut，提供了Error方法

#### PtocAttr
进程属性strcut，字段如下
* Dir String
* Env []String
* Files `[]*File`
* Sys `*syscall.SysProcAttr`

#### Process 
* FindProcess(pid int) 根据pid返回一个正在运行的进程的ProcessState，UNIX系统中始终返回，无论是否运行。
* StratProcess 
* p.Kill() 以立即退出的形式杀进程
* p.Release 释放进程相关的资源，使其不可用。只需要在Wait调用失败后使用
* p.Signa 发送信号量到进程
* P.Wait 等待进程退出，释放资源，返回 ProcessState 类型

#### ProcessState 进程状态，strcut
* Exited()
* Pid()
* String()
* Success()
* Sys()
* SysUsage()
* SystemTime()
* UserTime()

#### Signal 系统信号的抽象，接口类型
UNIX是syscall.Signal

#### 其他方法
* 环境变量
```
os.Expand() / os.ExpandEnv()
fmt.Println(os.Expand("$USER", os.Getenv))
```
* 用户和用户组
```
os.Getuid()
os.Geteuid()
os.Getgid()
os.Getegid()
os.Getgroups()
```
* 进程
* 文件目录权限等



