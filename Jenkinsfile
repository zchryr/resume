pipeline {
   agent any
   stages {
      stage('Build') {
        agent {
            docker {
               image 'blang/latex:ubuntu'
               args '-u root'
            }
         }
         steps {
            sh 'apt update; apt install fonts-font-awesome'
            sh 'pdflatex main.tex'
         }
      }
   }

   post {
      always {
         cleanWs()
         deleteDir()
      }
      cleanup{
         cleanWs()
         deleteDir()
      }
   }    
}