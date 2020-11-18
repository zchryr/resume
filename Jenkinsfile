pipeline {
   agent linux
   stages {
      stage('Build') {
        agent {
            docker {
               image 'blang/latex:ubuntu'
            }
         }
        steps {
            sh 'pdflatex -interaction nonstopmode main.tex'
        }
      }
   }
}