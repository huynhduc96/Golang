# System interface definition

## API Design

<!-- - Account Feature: Create account, Login, update Password, name.

- User Profile Feature: View My Profile, View Other User's Profile

- User Action: Post, Comment on Post, Follow/Unfollow other user. -->

### Account

- Sign up: **POST** `v1/api/account`

  Body : {Name, Username, Password, Email}

  Response : {Message}

- Login : **POST** `v1/api/account/login`

  Body : {Username, Password}

  Response : {Message}

- Edit Profile : **PUT** `v1/api/account`

  Body : {Name, Password, Email}

  Response : {Message}

### Profile

- View Follow list : **GET** `v1/api/friends/user_id`

  Response : {Follow List User}

- Follow New User : **POST** `v1/api/friends/user_id`

  Body : {User ID to follow}

  Response : {Message}

- Unfollow User : **DELETE** `v1/api/friends/user_id`

  Body : {User ID to unfollow}

  Response : {Message}

### Post

- View Posts : **GET** `v1/api/friends/user_id/posts`

  Response : {Post List}

- View Post : **GET** `v1/api/posts/post_id`

  Response : {Post}

- Create Post : **POST** `v1/api/posts`

  Body : {Post Content, Image}

  Response : {Message}

- Edit Post : **PUT** `v1/api/posts/post_id`

  Body : {Post Content, Image}

  Response : {Message}

- Delete Post : **DELETE** `v1/api/posts/post_id`

  Response : {Message}

### Comment

- View Comments : **GET** `v1/api/comments/post_id`

  Response : {Comment List}

- Create Comment : **POST** `v1/api/comments`

  Body : {Post ID, Content}

  Response : {Message}

### Like

- View Likes : **GET** `v1/api/likes/post_id`

  Response : {Like List}

- Create Like/Unlike : **POST** `v1/api/likes`

  Body : {Post ID, User ID}

### Newsfeed

- View Newsfeed : **GET** `v1/api/newsfeed`

Response : {Posts}
