pipeline {
    agent any
    // Install the desired Go version
    def root = tool name: '1.14', type: 'go'
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
                    sh 'go build ./cmd/beers-cli'
                }
                //sh 'go build ./cmd/beers-cli'
            }
        }
    }
}