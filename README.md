
# url-redirector
## simple program with:
* user creation
* user login
* adding urls
* activating urls (updating) 
* getting user's urls
* redirecting to users activated 1 url

#### Register

```http
  POST /auth/register
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required**, "email"|
| `name` | `string` | **Required**, "name"|
| `surname` | `string` | **Required**, "surname"|
| `password` | `string` | **Required**, "password"|



#### Login

```http
  GET /auth/login
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email` | `string` | **Required**, "email"|
| `password` | `string` | **Required**, "password"|


#### Add URL
- Add to your url list one more url
```http
  POST /urls
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `url` | `string` | **Required**, "url"|


### Redirect to URL
- Enter personal URL to redirect user's(user_id) activated URL

```http
  GET /urls/:user_id
```


#### Activate
- Activated url is the url where users will be redirected
```http
  POST /urls/activate
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `ID` | `int64` | **Required**, "id",  *url id to activate*|


#### Get user urls

```http
  POST /urls/activate
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `user_id` | `int64` | **Required**, "user_id"|



