pipeline {
    agent any
    tools {
        go 'go-1.18'
    }
    environment {
        GO111MODULE = 'on'
        GOOS = 'linux'
        GOARCH = 'amd64'
        CGO_ENABLED = '0' 
        DEVOPS_MAIL_PASSWORD = credentials("MAIL_PASSWORD")
    }
    stages {
        stage('Deploy'){
            steps {
                println('Deploying')
                script {
                    currentBuild.displayName = "mail"
                    if ( MAKEFILE == 'run' ) {
                        sh "make run"
                    } else {
                        error("Build Failed, ${MAKEFILE} is not defined")
                    }
                }
            }
        }
    }
    post {
        always {
            cleanWs()
        }
    }
}