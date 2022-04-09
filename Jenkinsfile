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
        MAIL_PASSWORD = credentials("MAIL_PASSWORD")
    }
    stages {
        stage('Deploy'){
            steps {
                println('Deploying')
                script {
                    currentBuild.displayName = "${SLS_ACTION}-${FUNCTION_NAME}"
                    if ( SLS_ACTION == 'run' ) {
                        sh "make run"
                    } else {
                        error("Build Failed, ${SLS_ACTION} is not defined")
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