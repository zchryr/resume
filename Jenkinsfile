pipeline {
   agent any
   stages {
      stage('Build') {
        agent {
            docker {
               image 'blang/latex:ctanfull'
            }
         }
        steps {
            sh 'xelatex main.tex'
        }
      }
   }
}