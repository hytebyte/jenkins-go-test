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
        sh '''go get -u github.com/jstemmer/go-junit-report
go test 2>&1 | go-junit-report > report.xml'''
        junit 'report.xml'
      }
    }
  }
  tools {
    go 'go-1.10'
  }
}