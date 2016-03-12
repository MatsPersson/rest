curl -H "Content-Type: application/json" -d "{\"task\":\"Learn golang\",\"duedate\":\"2016-03-11T11:45:26.371Z\", \"completed\": false}" http://127.0.0.1:8080/task/add -v
curl -H "Content-Type: application/json" -d "{\"task\":\"golang Rest server\",\"duedate\":\"2016-03-11T11:46:26.371Z\", \"completed\": false}" http://127.0.0.1:8080/task/add -v

curl -X GET http://127.0.0.1:8080/task/list

curl -X PUT -H "Content-Type: application/json" -d "{\"task\":\"Learn golang\",\"duedate\":\"2016-03-11T11:45:26.371Z\", \"completed\": true}" http://127.0.0.1:8080/task/1 -v

curl -X GET http://127.0.0.1:8080/task/1

curl -X GET http://127.0.0.1:8080/task/list

curl -X DELETE http://127.0.0.1:8080/task/delete/3

curl -X DELETE http://127.0.0.1:8080/task/delete/1

curl -X DELETE http://127.0.0.1:8080/task/delete/2

curl -X GET http://127.0.0.1:8080/task/list

curl -X GET http://127.0.0.1:8080/task/2


Expected (sample) output from above curl runs:

[D:\GoWorkspace\src\github.com\MatsPersson\rest]curl -H "Content-Type: application/json" -d "{\"task\":\"Learn golang\",\"duedate\":\"2016-03-11T11:45:26.371Z\", \"completed\": false}" http://127.0.0.1:8080/task/add -v
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> POST /task/add HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.47.1
> Accept: */*
> Content-Type: application/json
> Content-Length: 80
>
* upload completely sent off: 80 out of 80 bytes
< HTTP/1.1 201 Created
< Content-Type: application/json
< Date: Sat, 12 Mar 2016 03:16:45 GMT
< Content-Length: 78
<
{"task":"Learn golang","duedate":"2016-03-11T11:45:26.371Z","completed":false}* Connection #0 to host 127.0.0.1 left intact

[D:\GoWorkspace\src\github.com\MatsPersson\rest]curl -H "Content-Type: application/json" -d "{\"task\":\"golang Rest server\",\"duedate\":\"2016-03-11T11:46:26.371Z\", \"completed\": false}" http://127.0.0.1:8080/task/add -v
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> POST /task/add HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.47.1
> Accept: */*
> Content-Type: application/json
> Content-Length: 86
>
* upload completely sent off: 86 out of 86 bytes
< HTTP/1.1 201 Created
< Content-Type: application/json
< Date: Sat, 12 Mar 2016 03:16:45 GMT
< Content-Length: 84
<
{"task":"golang Rest server","duedate":"2016-03-11T11:46:26.371Z","completed":false}* Connection #0 to host 127.0.0.1 left intact

[D:\GoWorkspace\src\github.com\MatsPersson\rest]
[D:\GoWorkspace\src\github.com\MatsPersson\rest]curl -X GET http://127.0.0.1:8080/task/list
{"1":{"task":"Learn golang","duedate":"2016-03-11T11:45:26.371Z","completed":false},"2":{"task":"golang Rest server","duedate":"2016-03-11T11:46:26.371Z","completed":false}}
[D:\GoWorkspace\src\github.com\MatsPersson\rest]
[D:\GoWorkspace\src\github.com\MatsPersson\rest]curl -X PUT -H "Content-Type: application/json" -d "{\"task\":\"Learn golang\",\"duedate\":\"2016-03-11T11:45:26.371Z\", \"completed\": true}" http://127.0.0.1:8080/task/1 -v
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> PUT /task/1 HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.47.1
> Accept: */*
> Content-Type: application/json
> Content-Length: 79
>
* upload completely sent off: 79 out of 79 bytes
< HTTP/1.1 201 Created
< Content-Type: application/json
< Date: Sat, 12 Mar 2016 03:16:45 GMT
< Content-Length: 77
<
{"task":"Learn golang","duedate":"2016-03-11T11:45:26.371Z","completed":true}* Connection #0 to host 127.0.0.1 left intact

[D:\GoWorkspace\src\github.com\MatsPersson\rest]
[D:\GoWorkspace\src\github.com\MatsPersson\rest]curl -X GET http://127.0.0.1:8080/task/1
{"task":"Learn golang","duedate":"2016-03-11T11:45:26.371Z","completed":true}
[D:\GoWorkspace\src\github.com\MatsPersson\rest]
[D:\GoWorkspace\src\github.com\MatsPersson\rest]curl -X GET http://127.0.0.1:8080/task/list
{"1":{"task":"Learn golang","duedate":"2016-03-11T11:45:26.371Z","completed":true},"2":{"task":"golang Rest server","duedate":"2016-03-11T11:46:26.371Z","completed":false}}
[D:\GoWorkspace\src\github.com\MatsPersson\rest]
[D:\GoWorkspace\src\github.com\MatsPersson\rest]curl -X DELETE http://127.0.0.1:8080/task/delete/3

[D:\GoWorkspace\src\github.com\MatsPersson\rest]
[D:\GoWorkspace\src\github.com\MatsPersson\rest]curl -X DELETE http://127.0.0.1:8080/task/delete/1

[D:\GoWorkspace\src\github.com\MatsPersson\rest]
[D:\GoWorkspace\src\github.com\MatsPersson\rest]curl -X DELETE http://127.0.0.1:8080/task/delete/2

[D:\GoWorkspace\src\github.com\MatsPersson\rest]
[D:\GoWorkspace\src\github.com\MatsPersson\rest]curl -X GET http://127.0.0.1:8080/task/list
{}
[D:\GoWorkspace\src\github.com\MatsPersson\rest]
[D:\GoWorkspace\src\github.com\MatsPersson\rest]curl -X GET http://127.0.0.1:8080/task/2
Task not foundnull
[D:\GoWorkspace\src\github.com\MatsPersson\rest]