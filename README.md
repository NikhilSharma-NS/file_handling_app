# File Store APP


### Adding the file to the store 

Request : /store
Method: POST

Response :
```
Status Code 200 
Response filename
```
```
Status Code 500
Response errormsg
```

Sample Curl 

```
curl -X POST \
  'http://localhost:8080/store' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -H 'postman-token: b8922516-52bb-fd4f-2fe6-e4a8a5ebf4a9' \
  -F file=@file.txt
```

Sample Response 

```
file.txt
```

### Updating the file to the store 

Request : /store
Method: PATCH

Response :
```
Status Code 200 
Response filename
```
```
Status Code 500 
Response errormsg
```


Sample Curl :

```
curl -X PATCH \
  'http://localhost:8080/store' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -H 'postman-token: 797ec70d-9c2d-5af3-b9cc-f63202deb9b0' \
  -F file=@file.txt
```

Sample Response 

```
file.txt
```

### List all the file from store 

Request : /store
Method: GET

Response :

```
Status Code 200 
Response [filename]
```

```
Status Code 500 
Response errormsg
```


Sample Curl :
```
curl -X GET \
  'http://localhost:8080/store' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -H 'postman-token: 10459757-ad92-4c61-a4af-2ed8acde507d' \
```

Sample Response 

```
["file.txt"]
```

### Delete file from store

Request : /store?filename={file_name}
Method: DELETE

Response :
```
Status Code 200 
Response file_name
```
```
Status Code 500 errormsg
```

Sample Curl :

```
curl -X DELETE \
  'http://localhost:8080/store?filename=file.txt' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -H 'postman-token: cf51439e-798c-e981-e26e-522e6ae19106' \
```


Sample Response 
```
file.txt
```
