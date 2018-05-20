pipeline {
  agent any
  stages {
    stage('Install Go') {
      steps {
        tool(name: 'go-1.10', type: 'go')
      }
    }
    stage('Build') {
      steps {
        sh 'go version'
      }
    }
  }
}