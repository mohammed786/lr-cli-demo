# Introduction

This is the LoginRadius CLI project

## List of Commands supported in Beta

```
lr
    - help
    - register
    - login
    - get
        - servertime
        - config
        - social <provider>
        - domain
        - account
        - sites
        - theme
    - add
        - social <provider>
        - domain
        - account
    - delete
        - social <provider>
        - domain
        - account --uid, --email
    - set
        - social <provider>
        - domain
        - account --uid
        - account-email --uid
        - account-phone --uid
        - sec-ques --uid
        - theme
    - verify
        - email <email>
        - username <username>
        - resend <username>, <email>
        - invalidate 
    
    - get-password 
    - set-password
    - reset-secret
```   
       
## Setup Enviroment File

- Create `app.env` file in the root of the project.
- Below is the list of required ENV vairables

  ```
  LOGINRADIUS_API_KEY
  LOGINRADIUS_API_DOMAIN
  ADMINCONSOLE_API_DOMAIN
  ```