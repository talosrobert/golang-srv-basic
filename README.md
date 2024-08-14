# golang-srv-basic

golang REST server with basic authentication

## endpoints

POST   /ad/ `create an ad, returns ID`
GET    /ad/<adid> `returns a single ad by ID`
GET    /ad/ `returns all ads`
DELETE /ad/<adid> `delete an ad by ID`
GET    /tag/<tagname> `returns list of ads with this tag`
