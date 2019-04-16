def project = 'anishnath'
def  appName = 'hello'
def  imageTag = "${project}/${appName}:${env.BRANCH_NAME}.${env.BUILD_NUMBER}"
def  svcName = "${appName}"



pipeline {  
  
  environment {
    registry = "anishnath/hello"
    registryCredential = 'docker-hub-creds'
    dockerImage = ''
  }
  
  agent {
    kubernetes {
      label 'hello-app'
      defaultContainer 'jnlp'
      yaml """
apiVersion: v1
kind: Pod
metadata:
labels:
  component: ci
spec:
  # Use service account that can deploy to all namespaces
  serviceAccountName: jenkins
  containers:
  - name: golang
    image: golang:1.10
    command:
    - cat
    tty: true
  - name: kubectl
    image: gcr.io/cloud-builders/kubectl
    command:
    - cat
    tty: true
  - name: docker
    image: docker
    command:
    - cat
    tty: true  
    volumeMounts:
    - name: dockersock
      mountPath: "/var/run/docker.sock"
  volumes:
   - name: dockersock
     hostPath:
       path: /var/run/docker.sock
"""
}
  }
  stages {
    stage('Test') {
      steps {
        container('golang') {
          sh """
            ln -s `pwd` /go/src/hello-app
            cd /go/src/hello-app
            go test
          """
        }
      }
    }
    
    
    
    stage('Build and push image with Container Builder') {
      steps {
        container('docker') {
            //sh "which docker"
            //sh "docker build -t ${imageTag} ."
             script {
                dockerImage = docker.build "$imageTag"
                docker.withRegistry( '', registryCredential ) {
                dockerImage.push()
                  }
           } 
        }
      }
    }
    
    
    stage('Deploy Canary') {
      // Canary branch
      when { branch 'canary' }
      steps {
        container('kubectl') {
          sh("sed -i.bak 's#anishnath/hello:latest#${imageTag}#' ./canary/*.yaml")
          sh("kubectl --namespace=production apply -f canary/deploy.yaml")
          sh("echo http://`kubectl --namespace=production get service/kuebernetes-by-example-service -o jsonpath='{.status.loadBalancer.ingress[0].ip}'` > ${svcName}")
        } 
      }
    }
    stage('Deploy Production') {
      // Production branch
      when { branch 'master' }
      steps{
        container('kubectl') {
          sh("sed -i.bak 's#anishnath/hello:latest#${imageTag}#' ./production/*.yaml")
          sh("kubectl --namespace=production apply -f production/deploy.yaml")
          sh("echo http://`kubectl --namespace=production get service/kuebernetes-by-example-service -o jsonpath='{.status.loadBalancer.ingress[0].ip}'` > ${svcName}")
        }
      }
    }
    stage('Deploy Dev') {
      // Developer Branches
      when { 
        not { branch 'master' } 
        not { branch 'canary' }
      } 
      steps {
        container('kubectl') {
          // Don't use public load balancing for development branches
          // Create a Seperate name space for the Development Branch 
          sh("kubectl get ns ${env.BRANCH_NAME} || kubectl create ns ${env.BRANCH_NAME}")
          sh("sed -i.bak 's#anishnath/hello:latest#${imageTag}#' ./dev/*.yaml")
          sh("kubectl --namespace=${env.BRANCH_NAME} apply -f de/deploy.yaml")
          echo 'To access your environment run `kubectl proxy`'
          echo "Then access your service via http://localhost:8001/api/v1/proxy/namespaces/${env.BRANCH_NAME}/services/${feSvcName}:80/"
        }
      }     
    }
  }
}

