# hello-world-with-gin
Hello world REST API with Gin web framework.

Following steps requires you to have a programatic access to Github.
Using an SSH key might be an option. See the [official GitHub documentation](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account).

## Clonning the repo to your local machine and resetting it
```shell
git clone git@github.com:tomasdembelli/hello-world-with-gin.git
rm -rf go.mod go.sum   # These will be created in the next step
```

## Setting up the code repository
[Official reference](https://go.dev/doc/tutorial/create-module).
```shell
go mod init github.com/<your_github_handler>/hello-world-with-gin 
go mod tidy   # this will install all the dependencies (observe that the go.mod and go.sum are created)
go run .
curl -X 'GET' 'http://localhost:8080/'  # Confirm that the app is running
# Kill the application
## Push your code to Github
git add .
git commit -m 'Initial commit'
git push
```
