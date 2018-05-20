pipeline {
  agent any
  stages {
    stage('Version') {
      steps {
        sh 'go version'
      }
    }
    stage('Test') {
      steps {
        sh 'go test'
      }
    }
  }
  tools {
    go 'go-1.10'
  }
}