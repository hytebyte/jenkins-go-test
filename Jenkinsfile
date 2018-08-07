pipeline {
  agent any
  stages {
    stage('Version') {
      steps {
        sh 'go version'
      }
    }
    stage('Script') {
      when { changeset "script/**" }
      steps {
        sh '''cd script
go run *.go'''
      }
    }
    stage('Test') {
      steps {
        sh '''go get -u github.com/jstemmer/go-junit-report
go test -v 2>&1 | ~/go/bin/go-junit-report > report.xml'''
        junit 'report.xml'
      }
    }
    stage('Build') {
      steps {
        sh 'go build'
      }
    }
  }
  tools {
    go 'go-1.10'
  }
}
