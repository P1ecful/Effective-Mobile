# Go Restful API TimeTracker

## Features
The following functions are a set for creating this web APIs:
- Routing with [Fiber](https://github.com/gofiber/fiber)
- Migrations and Database work support with [gorm](https://gorm.io)

## Start Application
  - Clone this repository
  - Create a postgres database. 
  - Change Data in `config/config.env`
  - Run the application: `go run main.go`

## API Routes
| Path          | Method | Request/Exmaple               |  Desription                                           |                                    
| ------------- | ------ | ----------------------------- | ----------------------------------------------------- |
| user/create   | GET   | { "Name": "John", "Surname": "Smith", "Address": "green street","Passport Number": "1234567890" } | Add user |   
| user/delete   | GET   | { "ID": 0 }                    | Delete user                                           |     
| user/update   | GET   | { "ID": 0, "Name": "John" }    | Update user`s data                                    |      
| user/get-users| GET   | { "ID": 0, "Name": "John" }    | Get users by filters                                  | 
| user/get-time | GET   | { "ID": 0 }                    | Get user`s time per a task                            | 
| task/create   | GET   | { "Task": "Create makefile", "User ID": 0 } | Create task  
| task/start    | GET   | { "ID": 0 }                    | Start a task(when prompted, enter the task ID)        |     
| task/finish   | GET   | { "ID": 0 }                    | Finish a task(when prompted, enter the task ID)       |  
