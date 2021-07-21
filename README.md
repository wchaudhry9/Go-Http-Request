# Go-Http-Request
Practice making HTTP requests with Golang

# To Run
Make sure Go is installed on your system

go run goDocker.go

You can either use curl or the web browser to trigger the request

# Running with Docker
Install the Docker engine on your system

Use the Dockerfile in docker/ to build the Docker image (docker build --tag <imageName> .*)
  
* refers to the current directory where the Dockerfile is located along with the Go code
 
Use the command "docker run -p <IP:Port:Port> <imageName>" to run the container
  
You can either curl the container or try accessing the exposed port via your web browser

# Future Additions
k8s coming soon
