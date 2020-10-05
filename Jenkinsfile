pipeline {
    agent {
        docker {
            image 'golang:alpine'
            args '-p 3001:3000'
        }
    }
    environment {
        GO111MODULE= 'on'
        APP_NAME = 'PUNKAPI'
    }
    stages {
        stage('Build') {
            steps {
                sh 'go build -v  ./cmd/beers-cli'
            }
        }
        stage('Test') {
            steps {
                sh 'go test -v ./internal/cli/fetching/service_test.go'
            }
        }
        stage('Deliver') {
            steps {
                sh 'go build -v  ./cmd/beers-cli'
                sh 'go test -v  ./internal/cli/fetching/service_test.go'
                sh 'go run ./cmd/beers-cli/main.go'
                
            }
        }
    }
}
