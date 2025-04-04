namespace go api.test

struct HelloReq {
    1: string Name (api.query="name");
}

struct TestGetReq {
    1: string Name (api.query="name");
    2: string Age (api.query="age", api.vd="$!='0'; msg:'不能为0'");
}

struct HelloResp {
    1: string RespBody;
}

struct TestGetResp {
    1: string Message;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
    TestGetResp TestGetMethod(1: TestGetReq request) (api.get="/get");
}