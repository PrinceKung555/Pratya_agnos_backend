# Pratya_Cherbundit_agnos_backend
## How to deploy.

1.Open Docker.

2.Run ```docker-compose up --build```

3.Connect to database and create table.

```
CREATE TABLE public.logger (
	request varchar(255) NULL,
	response varchar(255) NULL,
	"method" varchar(10) NULL,
	code varchar(5) NULL,
	accesstime timestamp NULL
);
```

3.Use Postman or Curl to test API POST method http://localhost:80/api/strong_password_steps.

Body raw
```
{
"init_password": "aA1"
}
```

## How to run unit test.

1.Go to services directory.
```
cd services
```
2.Run go test
```
go test -cover
```
