pipeline {
  agent any
  stages {
    stage('Install Go') {
      steps {
        tool(name: 'go', type: 'go-1.10')
      }
    }
    stage('Build') {
      steps {
        sh 'go version'
      }
    }
  }
}