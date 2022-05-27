# postgr8
Go functions for doing simple admin stuff with postgres. 


## Usage
run make godoc to start the local go doc server
```
make godoc
```
The connect with your browser
http://localhost:6060/pkg/github.com/natemarks/postgr8/


## Testing

The deployments/ folder contains AWS CDK code to create an open test database. It also puts the credentials in secretsmanager. before running tests, set your aws credentials and install the aws python CDK. Then run:
```make db-create
```

Now that the  database exists you can run all of the tests:
```
make deploymenttest
```