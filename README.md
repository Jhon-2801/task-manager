
# Task Manager

The project designed for educational purposes to learn development technologies and practices. The application is packaged with Docker for easy deployment.


## Tech Stack

**Server:**   
- **Lenguage:** Golang
- **ORM:** Gorm
- **Framework:** Gin
- **Database:** MySQL
- **Authentication:** JWT
- **Containerization:** Docker


## API Reference

#### Get all tasks

```http
  GET /api/tasks
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token`   | `string` | **Required**. Your Token   |


#### Get all tasks user

```http
  GET /api/tasks/${id_user}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`   | `string` | **Required**. Your Token          |

#### Create tasks

```http
  POST /api/tasks/${id}
```

| Parameter    | Type     | Description                    |
| :--------    | :------- | :------------------------------|
| `name`       | `string` | **Required**. Name Tasks       |
| `description`| `string` | description Tasks              | 
| `due_date`   | `string` | **Required**. Date Tasks       |
| `user_id`   | `string`  | **Required**. UserId           |

#### UpDate tasks

```http
  POST /api/update/${id_task}
```

| Parameter    | Type     | Description                    |
| :--------    | :------- | :------------------------------|
| `name`       | `string` | **Required**. Name Tasks       |
| `description`| `string` | description Tasks              | 
| `due_date`   | `string` | **Required**. Date Tasks       |
| `user_id`    | `string` | **Required**. UserId           |
| `status`     | `string` | **Required**. Status Task      |




