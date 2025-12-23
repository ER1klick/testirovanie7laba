pipeline {
    agent any
    tools {
        go 'go1.21.5'
    }
    stages {
        stage('Checkout') {
            steps { checkout scm }
        }
        stage('Test') {
            steps {
                sh 'go test -v .'
            }
        }
    }
}