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
            sh 'pdflatex main.tex'
         }
      }
   }

   post {
      always {
         cleanWs{
            cleanWhenAborted(true)
            cleanWhenFailure(true)
            cleanWhenNotBuilt(false)
            cleanWhenSuccess(true)
            cleanWhenUnstable(true)
            deleteDirs(true)
            notFailBuild(true)
            disableDeferredWipeout(true)
         }
      }
   }    
}