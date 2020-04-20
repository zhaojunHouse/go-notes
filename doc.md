    外部:
    curl -X POST   http://192.168.157.147:10935/membership.default/v1/get-membership  -H 'Content-Type: application/json'   -H 'X-User-Id:5054485'
    内部：
    curl -X POST   http://192.168.157.147:10935/membership.internal.default/v1/get-membership  -H 'Content-Type: application/json'   -d '{"userID":1}'
   
