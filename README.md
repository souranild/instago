# instago
Instagram Backend API clone, in Golang
(Appointy Tech Task)

## Completion
- [x] User Creation
- [x] Get user by User ID
- [x] Post Creation
- [x] Get a Post by Post ID
- [ ] Get all posts by a User


## Run the App:
1. Install Golang
2. Open folder containing main.go
3. Run `go run main.go`
4. The app will be hosted at localhost:8080 by default

## API Reference

#### Create a user
```http
  POST /users
  
  payload = {
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
  
  payload = {
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

![Instago](https://user-images.githubusercontent.com/42074408/136666072-f06dbb02-cb1e-4b6a-bb67-c0bef47fb27d.png)


