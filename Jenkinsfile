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
            sh 'xelatex main.tex'
        }
      }
   }
}