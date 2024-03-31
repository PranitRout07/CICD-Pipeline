def successfulStages = []
def failedStages = []


pipeline {
    agent any
    stages {
        stage('FETCH CODE') {
            steps {
                git branch: 'main', url: 'https://github.com/PranitRout07/CICD-Pipeline.git'

            }
            post {
                always {
                    script {
                        switch (currentBuild.currentResult) {
                            case 'SUCCESS':
                                successfulStages.add("FETCH CODE")
                                break
                            default:
                                failedStages.add("FETCH CODE")
                        }
                    }
                }
            }
        }
        stage('TEST') {
            environment {
                goHome = tool 'go'
            }

            steps{
                withCredentials([string(credentialsId: 'API_TOKEN', variable: 'API_TOKEN')]) {
                    script {
                        dir('backend') {
                        
                            sh "${goHome}/bin/go test"
                            sh "${goHome}/bin/go test -coverprofile=coverage.out ./..."
                        }
                    }
                }
            }
            post {
                always {
                    script {
                        switch (currentBuild.currentResult) {
                            case 'SUCCESS':
                                successfulStages.add("TEST")
                                break
                            default:
                                failedStages.add("TEST")
                        }
                    }
                }
            }
        
    }
	stage('CODE ANALYSIS') {
    	environment {
        	scannerHome = tool 'Sonar'
    	}
    	steps {
        	withSonarQubeEnv('Sonar') {
            	sh "${scannerHome}/bin/sonar-scanner \
                	-Dsonar.projectKey=weather-tracker \
                	-Dsonar.projectName=weather-tracker \
                	-Dsonar.projectVersion=1.0 \
                	-Dsonar.sources=./backend"
        	}
    	     }
            post {
                always {
                    script {
                        switch (currentBuild.currentResult) {
                            case 'SUCCESS':
                                successfulStages.add("CODE ANALYSIS")
                                break
                            default:
                                failedStages.add("CODE ANALYSIS")
                        }
                    }
                }
            }
	}

        stage("QUALITY GATE") {
            steps {
                timeout(time: 1, unit: 'HOURS') {
                    waitForQualityGate abortPipeline: true
                }
            }
            post {
                always {
                    script {
                        switch (currentBuild.currentResult) {
                            case 'SUCCESS':
                                successfulStages.add("QUALITY GATE")
                                break
                            default:
                                failedStages.add("QUALITY GATE")
                        }
                    }
                }
            }
        }

        stage('BUILD BACKEND DOCKER IMAGE'){
            steps {
                sh 'docker build -t backend-weather-app ./backend/'
            }
            post {
                always {
                    script {
                        switch (currentBuild.currentResult) {
                            case 'SUCCESS':
                                successfulStages.add("BUILD BACKEND DOCKER IMAGE")
                                break
                            default:
                                failedStages.add("BUILD BACKEND DOCKER IMAGE")
                        }
                    }
                }
            }
        }
        stage('BUILD FRONTEND DOCKER IMAGE'){
            steps {
                sh 'docker build -t frontend-weather-app ./frontend/'
            }
            post {
                always {
                    script {
                        switch (currentBuild.currentResult) {
                            case 'SUCCESS':
                                successfulStages.add("BUILD FRONTEND DOCKER IMAGE")
                                break
                            default:
                                failedStages.add("BUILD FRONTEND DOCKER IMAGE")
                        }
                    }
                }
            }
        }
        stage('PUSH DOCKER IMAGES TO DOCKERHUB'){
            
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerPass', usernameVariable: 'dockerUser')]) {
                   sh "docker login -u ${env.dockerUser} -p ${env.dockerPass}"
                   sh "docker tag backend-weather-app ${env.dockerUser}/backend-weather-app:${BUILD_NUMBER}"
                   sh "docker tag frontend-weather-app ${env.dockerUser}/frontend-weather-app:${BUILD_NUMBER}"
                   sh "docker push ${env.dockerUser}/backend-weather-app:${BUILD_NUMBER}"
                   echo "Successfully Pushed Server Image to dockerhub "
                   sh "docker push ${env.dockerUser}/frontend-weather-app:${BUILD_NUMBER}"
                   echo "Successfully Pushed Frontend Image to dockerhub"
                }
            }
            post {
                always {
                    script {
                        switch (currentBuild.currentResult) {
                            case 'SUCCESS':
                                successfulStages.add("PUSH DOCKER IMAGES TO DOCKERHUB")
                                break
                            default:
                                failedStages.add("PUSH DOCKER IMAGES TO DOCKERHUB")
                        }
                    }
                }
            }
        }
         stage('TRIVY SCAN') {
             steps {
             	withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerPass', usernameVariable: 'dockerUser')]) {

                      sh "docker run --rm -v D:/trivy-report/:/root/.cache/ aquasec/trivy:0.18.3 image ${env.dockerUser}/go-server:latest > trivy-report-backend-${BUILD_NUMBER}.txt"
                      sh "docker run --rm -v D:/trivy-report/:/root/.cache/ aquasec/trivy:0.18.3 image ${env.dockerUser}/frontend:latest > trivy-report-frontend-${BUILD_NUMBER}.txt"
                }
 		    }
             

             post {
                 success {
                     script {
                        
                         archiveArtifacts artifacts: 'trivy-report-*.txt'
                        
                        
                         emailext (
                             subject: "Trivy Scan Report Summary",
                             body: "The pipeline completed successfully. Please find the Trivy reports attached.",
                             to: 'pranitrout72@gmail.com',
                             from: 'jenkins@example.com',
                             replyTo: 'jenkins@example.com',
                             mimeType: 'text/html',
                             attachmentsPattern: "trivy-report-*.txt"
                         )
                        
                 
                        successfulStages.add("TRIVY SCAN") 
                     }
                 }
                 always {
                     script {
                   
                         if (currentBuild.currentResult == 'FAILURE') {
                             failedStages.add("TRIVY SCAN")
                         }
                     }
                 }
             }

         }

        stage('UPDATE MANIFEST FILES') {
            environment {
                GIT_REPO_NAME = "Manifest-Files"
                GIT_USER_NAME = "PranitRout07"
                GIT_PASSWORD = credentials('github')
            }
            steps {
                script {
                    dir('Manifest-Files') {
                        withCredentials([string(credentialsId: 'github', variable: 'GIT_PASSWORD')]) {
                            echo "Cloning repository..."
                            sh "git config --global user.email 'pranitrout72@gmail.com'"
                            sh "git config --global user.name 'PranitRout07'"
                            sh "git init"
                            def remoteExists = sh(script: "git remote show origin", returnStatus: true)
                            if (remoteExists == 0) {
                                echo "Remote origin already exists"
                            } else {
                                echo "Adding remote origin..."
                                sh "git remote add origin https://${GIT_USER_NAME}:${env.GIT_PASSWORD}@github.com/${GIT_USER_NAME}/${GIT_REPO_NAME}.git"
                            }
                            sh "git fetch origin"
                            sh "git checkout main"
                            
                            echo "Updating deployment file..."
                            sh "sed -i 's/image: pranit007\\/frontend-weather-app:[0-9]\\+/image: pranit007\\/frontend-weather-app:${BUILD_NUMBER}/g' deployment.yml"
                            sh "sed -i 's/image: pranit007\\/backend-weather-app:[0-9]\\+/image: pranit007\\/backend-weather-app:${BUILD_NUMBER}/g' deployment.yml"
                    
                            sh "git add deployment.yml"
                            sh "git commit -m 'Updated deployment file to version ${BUILD_NUMBER}'"
                            
                            echo "Pushing changes to GitHub..."
                            sh "git push origin main"
                        }
                    }
                }
            }
        
        
        post {
            always {
                script {
                    switch (currentBuild.currentResult) {
                        case 'SUCCESS':
                            successfulStages.add("UPDATE MANIFEST FILES")
                            break
                        default:
                            failedStages.add("UPDATE MANIFEST FILES")
                    }
                }
            }
        }
    }

    }

    post {
        always {
            script {
                SuccessfulStageColor = 'green'
                FailureStageColor = 'red'
                def emailBody = """<html>
                                    <body>
                                        <b>Successful Stages:</b><br>
                                        <b><p style='color: ${SuccessfulStageColor};'>${generateList(successfulStages)}<br></p></b>
                                        
                                        <b>Failed Stages:</b><br>
                                        <b><p style='color: ${FailureStageColor};'>${generateList(failedStages)}<br></p></b>
                                        
                                    </body>
                                </html>"""
                
                emailext (
                    subject: "Pipeline Summary: ${currentBuild.currentResult}",
                    body: emailBody,
                    to: 'pranitrout72@gmail.com',
                    from: 'jenkins@example.com',
                    replyTo: 'jenkins@example.com',
                    mimeType: 'text/html'
                )
            }
        }
    }
}

def generateList(list) {
    def numberedText = ''
    list.eachWithIndex { item, index ->
        numberedText += "${index + 1}. ${item}<br>"
    }
    return numberedText
}   
