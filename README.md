# Pratya_Cherbundit_agnos_backend
How to deploy.

1.Run ```docker-compose up --build```

2.Use Postman or Curl to test API POST method http://localhost:80/api/strong_password_steps.

Body raw
```
{
"init_password": "aA1"
}
```

How to run unit test.

1.Go to services directory
```
cd services
```
2.Run go test
```
go test -cover
```
