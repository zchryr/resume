pipeline {
   agent {
      label 'docker'
   }
   
   parameters {
      booleanParam(name: 'release', defaultValue: false, description: 'Create a Gitea repo release.')
      booleanParam(name: 'upload', defaultValue: false, description: 'Upload to S3 bucket.')
   }

   environment {
      // Git stuff.
      GITEA_API_TOKEN = credentials('gitea-access-key')
      REPO_NAME = sh(script: "basename \$(git remote get-url origin) .git", returnStdout: true)
      REPO_OWNER = sh(script: "git config --get remote.origin.url | cut -d'/' -f4", returnStdout: true)

      // AWS stuff.
      AWS_DEFAULT_REGION = credentials('AWS_DEFAULT_REGION')
      AWS_ACCESS_KEY_ID = credentials('AWS_ACCESS_KEY_ID')
      AWS_SECRET_ACCESS_KEY = credentials('AWS_SECRET_ACCESS_KEY')

      // Docker stuff.
      registryRepo = 'registry.rohrbach.xyz/python-runtime:latest'
      registryAddress = 'https://registry.rohrbach.xyz'
      registryCredential = 'rohrbach-registry'
   }

   stages {
      stage('Build LaTeX Resume') {
         steps {
            agent { 
               label ("ubuntu") 
            }
            script {
               docker.image('blang/latex:ctanfull').inside {
                  sh 'pdflatex main.tex'
                  sh 'mv main.pdf Zachary-Rohrbach-Resume.pdf'
               }
            }
         }
      }
      stage('Create Gitea Release'){
         when { expression { params.release } }
         agent {
            docker { 
               image registryRepo
               registryUrl registryAddress
               registryCredentialsId registryCredential
               args '-u root:root'
            }
         }
         steps {
            sh "pip3 install -r ./scripts/requirements.txt -q"
            sh "python3 ./scripts/gitea-release.py"
         }
      }
      stage('Upload To S3'){
         when { expression { params.upload } }
         agent {
            docker {
               image registryRepo
               registryUrl registryAddress
               registryCredentialsId registryCredential
               args "-u root:root -v $WORKSPACE:/app"
            }
         }
         steps {
            sh "pip3 install -r ./scripts/requirements.txt -q"
            sh "cd /app && python3 ./scripts/upload-to-s3.py"
         }
      }
   }

   post {
      always {
         deleteDir()
      }
      success {
         mattermostSend color: 'good', message: "Build Number: $BUILD_NUMBER\nJob Name: $JOB_NAME\nBuild URL: $BUILD_URL", text: "$JOB_NAME Pipeline Passing :)"
      }
      failure {
         mattermostSend color: 'bad', message: "Build Number: $BUILD_NUMBER\nJob Name: $JOB_NAME\nBuild URL: $BUILD_URL", text: "$JOB_NAME Pipeline Failing :("
      }
   }
}