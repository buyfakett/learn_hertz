namespace go api.test

struct HelloReq {
    1: string Name (api.query="name");
}

struct TestGetReq {
    1: string Name (api.query="name");
    2: string Age (api.query="age", api.vd="$!='0'; msg:'年龄不能为0'");
}

struct TestPutReq {
    1: string Name (api.form="name");
    2: string Age (api.form="age", api.vd="$!='0'; msg:'年龄不能为0'");
}

struct HelloResp {
    1: string RespBody;
}

struct TestGetResp {
    1: string Message;
}

struct TestPutResp {
    1: string Message;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
    TestGetResp TestGetMethod(1: TestGetReq request) (api.get="/get");
    TestPutResp TestPostMethod(1: TestPutReq request) (api.put="/add");
}