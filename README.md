# goTest
GET /blacklist  - all emails response 
GET /blacklist/email@email.em - is or not in blacklist
POST /blacklist - add JSON {"email": "@email.at", "status": "string", "clientID": "1"} 
DELETE /blacklist/email@email.em - delete from blacklist
