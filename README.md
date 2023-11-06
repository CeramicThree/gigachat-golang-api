# gigachat-golang-api
Module for using GIGACHAT AI in Go 

## Set your's credentials and access token in .env file
For example:

```
baseUrl=BASE_URL
authUrl=AUTH_URL
credentials=YOUR_CREDENTIALS
access_token=YOUR_ACCESS_TOKEN
```

Also, you can set your api configuration in .env file:
```
model=GigaChat:latest
temperature=0.87
top_p=0.5
n=1
max_tokens=512
repetition_penalty=1.07
stream=false
update_interval=0
```

The usage example is located in the `./cmd/main.go` file.
