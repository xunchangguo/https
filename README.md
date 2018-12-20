# https
https static file server


生成SA https://blog.csdn.net/wangshubo1989/article/details/77508738
生成pem和key
openssl req -new -nodes -x509 -out server.pem -keyout server.key -days 3650 

https.exe -keyFile server.key -certFile server.pem -dir E:/uiwork/dist/static/ -listen ":8080"
