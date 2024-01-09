# gdns Open API

## A 记录

### 查询A记录

```
Method: GET

Path: /a/:domain 

QueryString Parameter: None

Request Body: None

Response Body:

{
	"domain":"xxx",
	"type": "A",
	"records": ["192.168.1.1", "192.158.1.2"]
}

```

### 创建A记录

```
Method: POST

Path: /a/:domain 

QueryString Parameter: None

Request Body:

{
	"records": ["192.168.1.1", "192.158.1.2"]
}

Response Body:

{
	"message":"ok"
}

```

### 更新A记录

```
Method: PUT

Path: /a/:domain 

QueryString Parameter: None

Request Body:

{
	"records": ["192.168.1.1", "192.158.1.2"]
}

Response Body:

{
	"message":"ok"
}

```

### 删除A记录

```
Method: DELETE

Path: /a/:domain 

QueryString Parameter: None

Request Body: None

Response Body:

{
	"message":"ok"
}

```
