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
            sh 'sudo apt-get install -y fonts-font-awesome'
            sh 'xelatex main.tex'
        }
      }
   }
}