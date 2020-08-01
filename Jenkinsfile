pipeline {
    agent {
        docker {
            image 'golang:alpine'
            args '-p 3001:3000'
        }
    }
    environment {
        CI = 'true'
        GO111MODULE=on 
        CGO_ENABLED=0 
        GOOS=linux 
        GOARCH=amd64
        GOCMD=go
        GORUN=$(GOCMD) run
        GOBUILD=$(GOCMD) build
        GOCLEAN=$(GOCMD) clean
        GOTEST=$(GOCMD) test
        GOMOD=$(GOCMD) mod
        APP_NAME=PUNKAPI
    }
    stages {
        stage('Build') {
            steps {
                sh '@$(GOBUILD) -o $(APP_NAME)/cmd/beers-cli -v  '
            }
        }
        stage('Test') {
            steps {
                sh '@$(GOTEST) $(APP_NAME)//internal/cli/fetching -v -cover'
            }
        }
        stage('Deliver') {
            steps {
                sh '@$(GOBUILD) -o $(APP_NAME)/cmd/beers-cli -v  '
                sh '@$(GOTEST) $(APP_NAME)/internal/cli/fetching -v -cover'
                sh '@$(GORUN) $(APP_NAME)/cmd/beers-cli/main.go'
                
            }
        }
    }
}
