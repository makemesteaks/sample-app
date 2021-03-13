
# Welcome to sample app!  
  
This app will output a string that is passed to the url. :tada:

`curl http://localhost:3000/helloworld`  
  
`{"message":"helloworld"}`  

## Running the app

This app assumes you are running `OSX` with `homebrew` and `go` installed.  

### Locally  
  
To use the app locally, just run:  
  
`go run main.go`
  
### Minikube 

#### Prerequisites  
    
The following software is required for a full deployment.
You should install each component by running the commands provided.
  
##### Minikube  
    
```  
brew install minikube hyperkit
minikube start --driver=hyperkit  
minikube addons enable ingress  
```  
  
##### Helm  
  
```  
brew install helm  
```
  
##### Nginx Ingress  
  
```  
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx  
helm repo update  
helm install ingress-nginx ingress-nginx/ingress-nginx  
```  

#### Run  

After all the previous software has been installed, just run:  
  
`./setup.sh`  
