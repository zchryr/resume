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
            sh 'apt update; apt install fonts-font-awesome -y'
         }
         steps {
            sh 'pdflatex main.tex'
         }
      }
   }

   post {
      always {
         cleanWs()
      }
   }    
}