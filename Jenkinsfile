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
                // Export environment variables pointing to the directory where Go was installed
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                    sh 'go version'
                }
                //sh 'go build ./cmd/beers-cli'
            }
        }
    }
}