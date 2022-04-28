# ReSTful (Representational State Transfer)

CRUD {Create, Read, Update, Destroy}

HTTP Methods => {GET, POST, PATCH, DELETE, OPTIONS, ...}

## Employee Management Server (JSON)

CRUD            Action        HTTP Method              URL                   request body              response body
---------------------------------------------------------------------------------------------------------------------
Read            Index         GET                   /employees                    -                     [{...}, ...]
Read            Show          GET                   /employees/{id}               -                     {...}
Create          Create        POST                  /employees                  {...}                   {id: , ...}
Update          Update        PUT                   /employees/{id}             {...}                   {...}
Update          Update        PATCH                 /employees/{id}             {some attrs}            {...}
Destroy         Destroy       DELETE                /employees/{id}               -                     - / {...}
