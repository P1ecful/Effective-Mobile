# Go Restful API ApplyToCourier

## Features
The following functions are a set for creating this web APIs:
- Routing with [Fiber](https://github.com/gofiber/fiber)
- Migrations and Database work support with [gorm](https://gorm.io)

## Start Application
  - Clone this repository
  - Create a postgres database. 
  - Change Data in `config/local.env`
  - Run the application: `go run main.go`

## API Routes
| Path          | Method | Request                       |  Desription                                           |                                    
| ------------- | ------ | ----------------------------- | ----------------------------------------------------- |
| user/create       | GET   |  { "Creator Id": 0, "Item category": "", "Item weight": "", "First Address Phone": "", "Second Address Phone": "", "First Address": { "Street": "",  "Home": 0, "Housing": 0, "Entrance": 0,"Floor": 0, "Flat": 0, "Intercom Code": ""}, "Second Address": { "Street": "",  "Home": 0, "Housing": 0,  "Entrance": 0, "Floor": 0,  "Flat": 0, "Intercom Code": ""}}                             | Add user |   
| user/delete       | GET   | { "Order Id": 0 }             | Delete user by order ID                                    |     
| user/update  | GET   | { "Creator Id": 0 }           | Update users data                               |                                |

###### In development...