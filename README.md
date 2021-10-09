# instago
Instagram Backend API clone, in Golang

## API Reference

#### Create a user
```http
  POST /users
  
  headers = {
    id
    name
    email
    password
  }
```

#### Get users By ID 

```http
  GET /users/{id}
```
#### Create a post

```http
  POST /posts
  
  headers = {
    id
    caption
    imgurl
    timestamp
  }
```
#### Get post By ID 

```http
  GET /posts/{id}
```
#### Get all posts by a user

```http
  GET /posts/users/{id}
```

