# To-Do List API
This is a To-Do List API

## Version: 1.0


### Security
****  

| bearerauth | *Bearer Auth* |
|------------|---------------|
| In         | header        |
| Name       | Authorization |

### /api/auth

#### POST
##### Summary:

Login with Email & Password

##### Description:

authenticate user

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| email | body | email address of the user | Yes | string |
| password | body | password of the user | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [domains.LoginResponse](#domains.LoginResponse) |
| 400 | Bad Request | [domains.ErrorResponse](#domains.ErrorResponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domains.ErrorResponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domains.ErrorResponse) |

#### PUT
##### Summary:

Refresh Authentication

##### Description:

Generating new access token using a refresh token. Only valid refresh token will generate new

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| refreshToken | body | refresh token possessed by the user | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [domains.RefreshAuthnResponse](#domains.RefreshAuthnResponse) |
| 400 | Bad Request | [domains.ErrorResponse](#domains.ErrorResponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domains.ErrorResponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domains.ErrorResponse) |

### /api/auth/register

#### POST
##### Summary:

Register A User

##### Description:

New user must have a unique email address

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| email | body | email address of the new user, must be unique | Yes | string |
| password | body | password of the new user | Yes | string |
| name | body | name of the new user | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [domains.SuccessResponse](#domains.SuccessResponse) |
| 409 | Conflict | [domains.ErrorResponse](#domains.ErrorResponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domains.ErrorResponse) |

### /api/tasks

#### GET
##### Summary:

Fetch Tasks

##### Description:

Fetch Tasks By Owner ID. Only valid users may have tasks

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer Token | Yes | string |
| page | query | page number, acting as offset | Yes | string |
| size | query | page size, acting as limit | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [domains.GetTaskResponse](#domains.GetTaskResponse) |
| 400 | Bad Request | [domains.ErrorResponse](#domains.ErrorResponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domains.ErrorResponse) |
| 404 | Not Found | [domains.ErrorResponse](#domains.ErrorResponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domains.ErrorResponse) |

#### POST
##### Summary:

Add A New Task To DB

##### Description:

Add A New Task

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer Token | Yes | string |
| title | body | task's title | Yes | string |
| password | body | task's description | Yes | string |
| isCompleted | body | whether the task is completed | Yes | boolean |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [domains.AddTaskResponse](#domains.AddTaskResponse) |
| 400 | Bad Request | [domains.ErrorResponse](#domains.ErrorResponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domains.ErrorResponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domains.ErrorResponse) |

### /api/tasks/{id}

#### DELETE
##### Summary:

Delete Task

##### Description:

Delete Tasks By ID. Only valid task may be deleted

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer Token | Yes | string |
| id | path | Task ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [domains.GetTaskByIdResponse](#domains.GetTaskByIdResponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domains.ErrorResponse) |
| 403 | Forbidden | [domains.ErrorResponse](#domains.ErrorResponse) |
| 404 | Not Found | [domains.ErrorResponse](#domains.ErrorResponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domains.ErrorResponse) |

#### GET
##### Summary:

Fetch Task

##### Description:

Fetch Tasks By ID. Only valid task may get returned

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer Token | Yes | string |
| id | path | Task ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [domains.GetTaskByIdResponse](#domains.GetTaskByIdResponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domains.ErrorResponse) |
| 403 | Forbidden | [domains.ErrorResponse](#domains.ErrorResponse) |
| 404 | Not Found | [domains.ErrorResponse](#domains.ErrorResponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domains.ErrorResponse) |

#### PUT
##### Summary:

Edit Task

##### Description:

Edit Tasks By ID. Only valid task may be edited

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer Token | Yes | string |
| id | path | Task ID | Yes | string |
| title | body | Title of the task | Yes | string |
| description | body | Description of the task | Yes | string |
| isCompleted | body | whether the task is completed | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [domains.GetTaskByIdResponse](#domains.GetTaskByIdResponse) |
| 400 | Bad Request | [domains.ErrorResponse](#domains.ErrorResponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domains.ErrorResponse) |
| 403 | Forbidden | [domains.ErrorResponse](#domains.ErrorResponse) |
| 404 | Not Found | [domains.ErrorResponse](#domains.ErrorResponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domains.ErrorResponse) |

### /api/tasks/{id}/mark

#### PUT
##### Summary:

Mark Task as Completed

##### Description:

Mark Task as Completed By ID. Only valid task may be marked

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer Token | Yes | string |
| id | path | Task ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [domains.GetTaskByIdResponse](#domains.GetTaskByIdResponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domains.ErrorResponse) |
| 403 | Forbidden | [domains.ErrorResponse](#domains.ErrorResponse) |
| 404 | Not Found | [domains.ErrorResponse](#domains.ErrorResponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domains.ErrorResponse) |

### /api/tasks/completed

#### GET
##### Summary:

Retrieve Completed Tasks

##### Description:

Retrieve Completed Tasks. Only authorized users may see their own tasks

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer Token | Yes | string |
| page | query | page number, acting as offset | Yes | string |
| size | query | page size, acting as limit | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [domains.GetTaskResponse](#domains.GetTaskResponse) |
| 401 | Unauthorized | [domains.ErrorResponse](#domains.ErrorResponse) |
| 500 | Internal Server Error | [domains.ErrorResponse](#domains.ErrorResponse) |

### Models


#### domains.AddTaskResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data | [domains.AddTaskResponseData](#domains.AddTaskResponseData) |  | No |
| message | string |  | No |

#### domains.AddTaskResponseData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | No |

#### domains.ErrorResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |

#### domains.GetTaskByIdResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data | [domains.Task](#domains.Task) |  | No |
| message | string |  | No |

#### domains.GetTaskResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data | [ [domains.Task](#domains.Task) ] |  | No |
| message | string |  | No |

#### domains.LoginResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data | [domains.LoginResponseData](#domains.LoginResponseData) |  | No |
| message | string |  | No |

#### domains.LoginResponseData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| accessToken | string |  | No |
| refreshToken | string |  | No |

#### domains.RefreshAuthnData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| accessToken | string |  | No |
| refreshToken | string |  | No |

#### domains.RefreshAuthnResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data | [domains.RefreshAuthnData](#domains.RefreshAuthnData) |  | No |
| message | string |  | No |

#### domains.SuccessResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |

#### domains.Task

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| createdAt | string |  | No |
| description | string |  | Yes |
| id | string |  | No |
| isCompleted | boolean |  | No |
| owner | string |  | No |
| title | string |  | Yes |
| updatedAt | string |  | No |
