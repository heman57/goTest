# goTest

**GET** /blacklist  - all emails response 

**GET** /blacklist/email@email.em - is or not in blacklist

**POST** /blacklist - add JSON `{"email": "@email.at", "status": "string", "clientID": "1"}` 

**PUT** /blacklist -  add multiply JSON `{"emails": ["@email.at"], "status": "string", "clientID": "1"} `

**DELETE** /blacklist/email@email.em - delete from blacklist

Clientid:{string} have to be in header for all request. POST and PUT using JSON clientID when creating db record.

