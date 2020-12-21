pipeline {
   agent any
   
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
      registryRepo = 'registry.rohrbach.xyz/python-runtime'
      registryAddress = 'https://registry.rohrbach.xyz'
      registryCredential = 'rohrbach-registry'
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
      stage('Create Gitea Release'){
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
            sh "python3 ./scripts/gitea-release.py -release ${params.release}"
         }
      }
      stage('Upload To S3'){
         agent {
            docker { 
               image registryRepo
               registryUrl registryAddress
               registryCredentialsId registryCredential
               args '-u root:root'
            }
         }
         script {
            docker.image(registryRepo).inside {
               sh "pip3 install -r ./scripts/requirements.txt -q"
               sh "python3 ./scripts/upload-to-s3.py -upload ${params.upload}"
            }
         }
         // steps {
         //    sh "pip3 install -r ./scripts/requirements.txt -q"
         //    sh "python3 ./scripts/upload-to-s3.py -upload ${params.upload}"
         // }
      }
   }
}