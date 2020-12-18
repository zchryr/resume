pipeline {
   agent any
   
   parameters {
      booleanParam(name: 'release', defaultValue: false, description: 'Create a Gitea repo release.')
      booleanParam(name: 'upload', defaultValue: false, description: 'Upload to S3 bucket.')
   }

   environment {
      GITEA_API_TOKEN = credentials('gitea-access-key')
      REPO_NAME = sh(script: "basename \$(git remote get-url origin) .git", returnStdout: true)
      REPO_OWNER = sh(script: "git config --get remote.origin.url | cut -d'/' -f4", returnStdout: true)
      AWS_DEFAULT_REGION = credentials('AWS_DEFAULT_REGION')
      AWS_ACCESS_KEY_ID = credentials('AWS_ACCESS_KEY_ID')
      AWS_SECRET_ACCESS_KEY = credentials('AWS_SECRET_ACCESS_KEY')
   }

   stages {
      stage('Build LaTeX Resume') {
         steps {
            script {
               docker.image('blang/latex:ctanfull').inside {
                  sh 'pdflatex main.tex'
                  sh 'mv main.pdf Zachary-Rohrbach-Resume.pdf'
               }
            }
         }
      }
      stage('Install Python Packages') {
         steps {
            sh "pip3 install -r ./scripts/requirements.txt"
         }
      }
      stage('Create Gitea Release'){
         steps {
            sh "python3 ./scripts/gitea-release.py -release ${params.release}"
         }
      }
      stage('Upload To S3'){
         steps {
            sh "python3 ./scripts/upload-to-s3.py -upload ${params.upload}"
         }
      }
   }
}