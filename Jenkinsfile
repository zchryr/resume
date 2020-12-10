pipeline {
   agent any
   stages {
      stage('Build') {
        agent {
            docker {
               image 'blang/latex:ubuntu'
            }
         }
         steps {
            sh 'sudo apt update; sudo apt install fonts-font-awesome'
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