## flag 命令行参数解析

###类型

#### Flag
* 属性值Name(string), Usage(string), Value(Value), Default(string)

#### FlagSet：一组Flag
* 属性：Usage(func())  用于解析参数时的错误处理。
方法:
* Bool：参数为name（参数名）, value（参数默认值）, usage（参数用法）, 返回这个参数的值的内存地址
* BoolVar 参数为p（参数指针）, name, value, usage，无返回值，参数的内存地址存于p
* Duration/DurationVar
* Float64/Float64Var
* Int/IntVar/Int64/Int64Var
* String/StringVar
* Unit/UnitVar/Unit64/Unit64Var
* Var: 参数为value（默认类型和默认值）,name（参数名）, usage（用法）无返回值
```
authHTTPAddresses := app.StringArray{}
flagSet.Var(&authHTTPAddresses, "auth-http-address", "<addr>:<port> to query auth server (may be given multiple times)")
```



### 模块函数
* Lookup 参数为Name，返回Name对应的Flag
