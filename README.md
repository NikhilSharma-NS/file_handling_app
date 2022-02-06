# file_handling_app


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

#### Steps for Testing the app on local (I have used below steps )

Step 1 : 

```

docker build -t nkhlsharma050/go-filestore-app .

docker run -it -p 3000:3000 nkhlsharma050/go-filestore-app

docker push nkhlsharma050/go-filestore-app

kubectl apply -f .\k8s\k8s_deploy.yaml

kubectl port-forward go-filestore-app-6676b8d877-qwkdf 8080:8080

curl -X POST \
  'http://localhost:8080/store' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -H 'postman-token: 969ef2ed-a9ef-5dfd-07b6-b560674fbf6a' \
  -F file=@file.txt

```

Step 2 : 

```

kubectl apply -f .\k8s\k8s_service.yaml


```
take the port number of svc 

```
kubectl get svc 
```
```
go-filestore-app-service   LoadBalancer   10.110.40.38   localhost     80:30753/TCP   17m

```
Request : 

```
curl -X GET \
  http://localhost:80/store \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -H 'postman-token: 5d382e9c-acac-ac67-5d6d-4ad309ef3a8a' \
  -F file=@file.txt

```

Response :

```
["file.txt"]
```


#### Pods and services view 

```
kubectl get all
```

```
NAME                                    READY   STATUS    RESTARTS   AGE
pod/go-filestore-app-6676b8d877-qwkdf   1/1     Running   0          38m

NAME                               TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
service/go-filestore-app-service   LoadBalancer   10.110.40.38   localhost     80:30753/TCP   29m
service/kubernetes                 ClusterIP      10.96.0.1      <none>        443/TCP        16h

NAME                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/go-filestore-app   1/1     1            1           72m

NAME                                          DESIRED   CURRENT   READY   AGE
replicaset.apps/go-filestore-app-6676b8d877   1         1         1       54m

```


#### Building the app on env

Step 1: 

```
kubectl apply -f .\k8s\k8s_deploy.yaml
```
Step 2 : 

```
kubectl apply -f .\k8s\k8s_service.yaml
```

Step 3 : Take the port number of svc and access the app 

```
kubectl get svc 
```
```
go-filestore-app-service   LoadBalancer   10.110.40.38   localhost     80:30753/TCP   17m

```
Request : 

```
curl -X GET \
  http://localhost:80/store \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -H 'postman-token: 5d382e9c-acac-ac67-5d6d-4ad309ef3a8a' \
  -F file=@file.txt
```