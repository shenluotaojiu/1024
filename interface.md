# service包下
## `GetData(DataRequest_User) returns (DataResponse_User) {}`
通过精确的账户名，来获取数据返回
```
    message DataRequest_User{  
        string requestName = 1;  
    }
```
使用者只需要传入DataRequest_User结构体下的requestName类型，就可以得到DataResponse_User结构体返回  
