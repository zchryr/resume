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
            sh 'whoami'
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