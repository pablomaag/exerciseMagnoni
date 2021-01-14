# exerciseMagnoni

Topic: 

Develop a golang based simple web service that lists the files in a provided user’s home folder.
Example:  http://localhost:port/list/:userId returns the list of files and folders in userIds home directory

    Only HTTP GET is supported. All other HTTP methods should return an appropriate HTTP error code
    Invoking /list without a userId returns simple instructions
    If there is no home directory for the provided user, return a “not found” error using the appropriate HTTP error code
    A valid request returns a listing of the folders contents (files and folders) (preferably in a JSON format, not required)
    
Development:

I used GoLand as IDE and Postman to do the requests. 
Also, i used a library called Gorilla in order to handle the pathvariable "userId".
This library, can be cloned from its own repository in the GOROOT path -> git clone git://github.com/gorilla/mux.git
