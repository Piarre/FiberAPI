# Fiber API
Fiber API built with Go

### ⚙️ Setup
Install [AIR](https://github.com/cosmtrek/air) to start the server.

### 🖐 Endpoint

`/api/v1/users`


|          | Get all users | Find a user by ID  | Create a user | Delete a user | Update a user |
| :------: | :-----------: | :----------: | :-----------: | :-----------: | :-----------: |
|  Method  |     `GET`     |    `GET`     |    `POST`     |   `DELETE`  |    `PUT`     |
| Endpoint |    `/`     | `/:id` |    `/`     |  `/:id`  | `/:id` |
|   Body   |               |          |      ```{"first_name": "Piarre", "last_name": "Dev"} ``` | | ```{"first_name": "Piarre_", "last_name": "Dev_"}``` | 

-------------

## 🖐 Use
```bash
git clone https://github.com/Piarre/FiberAPI
air
```

# 🔐 License
MIT
