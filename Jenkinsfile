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