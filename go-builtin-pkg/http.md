#### RoundTripper
表示执行单个http事务的接口。根据Request获取Response。是go http模块执行http请求的底层，
不应该执行解释response，处理跳转等事务，只负责发送请求获取响应。
接口方法：RoundTrip(*Request) (*Response, error)
#### Transport 
RoundTripper 的实现
属性字段
* Proxy：参数为`*Request`，返回值为`*url.URL`和error 的函数。用于设置代理
* DialContext：
* Dial
* DialTLS
* TLSClientConfig
* TLSHandshakeTimeout
* DisableKeepAlives：布尔型，设置是否keep-alive
* DisableCompression：
* MaxIdelConns：int 最大空闲持久连接数，即处于空闲状态的keep-alive连接数
* MaxIdleConnsPerHost：int 
* IdleConnTimeout：time.Duration，超时
* ResponseHeaderTimeout
* ExpectContinueTimeout
* TLSNextProto
* MaxResponseHeaderBytes：
方法
