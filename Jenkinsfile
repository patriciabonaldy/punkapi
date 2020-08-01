pipeline {
    agent {
        docker {
            image 'golang:alpine'
            args '-p 3001:3000'
        }
    }
    environment {
        GO111MODULE= 'on'
        APP_NAME=PUNKAPI
    }
    stages {
        stage('Build') {
            steps {
                sh 'go build -v  $(APP_NAME)/cmd/beers-cli'
            }
        }
        stage('Test') {
            steps {
                sh 'go test -v $(APP_NAME)/internal/cli/fetching'
            }
        }
        stage('Deliver') {
            steps {
                sh 'go build -v  $(APP_NAME)/cmd/beers-cli'
                sh 'go test -v $(APP_NAME)/internal/cli/fetching'
                sh 'go run $(APP_NAME)/cmd/beers-cli/main.go'
                
            }
        }
    }
}
