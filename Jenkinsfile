pipeline {
    agent any
    tools {
        go '1.14'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'go build ./cmd/beers-cli'
            }
        }
    }
}