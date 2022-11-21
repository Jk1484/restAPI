pipeline {
  // Run on an agent where we want to use Go
  agent any

  // Ensure the desired Go version is installed for all stages,
  // using the name defined in the Global Tool Configuration
  tools {
    go "1.19.3"
  }

  stages {
    stage('Build') {
      agent any
      steps {
        // Output will be something like "go version go1.19 darwin/arm64"
        sh 'ls'
        sh 'docker compose up --build -d'
      }
    }
  }
}