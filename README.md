# My first terraform provider
This is a project made only to gain experience in developing a terraform provider. Also I am using this to improve my Golang programming skills

## Objectives
 - Develop simple API to save numbers and strings
 - Develop a terraform provider to control the resources via terraform


## Details on the API

### Summary
Developed in Golang it will provide some endpoints to later interact with them via terraform.
No pagination is implemented, we are assuming maximun of 10 elements at memory. No authentication is implemented

### Data structure
The data will be saved on memory and it will be needed to re-initialize it every time the API runs. There is no point on setting up a database as this project is for learning/practicing propouses.
Each element will contain the following structure (provided one element as example):

```json
{
    "id": "8c9da44dd0c38dd2f6e1f9cb288bd6e8", // MD5 hash of the creation timestamp with timezone included
    "number": 25,
    "string_data": "Velez Campeon 2024"
}
```
Empty strings are not allowed and number 0 are not allowed

### Updating data contitions
Data updates will be allowed only if the number is to be change. If the user requires to change the string_data, needs to delete the element and create it back again

### Api endpoints
These are the API endpoints that are implemented
 - **/all_data:** Making GET request will return all the elements saved on memory on the form of a list
 - **/id/\<element-id>:** Making GET request will return the selected element
 - **/create:** Making POST including in body a json with the number and the string, it will return a json with the recently created element
 - **/update/\<element-id>:** Making POST including in body a json with the number, it will return a json with the recently updated element
 - **/delete/\<element-id>:** Making GET request will delete the selected element