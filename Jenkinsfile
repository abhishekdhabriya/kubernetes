def project = 'anishnath'
def  appName = 'hello'
def  imageTag = "${project}/${appName}:${env.BRANCH_NAME}.${env.BUILD_NUMBER}"

podTemplate(label: 'kube-by-example', 
  containers: [
    containerTemplate(
      name: 'jnlp',
      image: 'jenkinsci/jnlp-slave:3.10-1-alpine',
      args: '${computer.jnlpmac} ${computer.name}'
    ),
    containerTemplate(
      name: 'alpine',
      image: 'twistian/alpine:latest',
      command: 'cat',
      ttyEnabled: true
    ),
  ],
  volumes: [ 
    hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock'), 
  ]
)
{

pipeline {
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
  - name: alpine
    image: twistian/alpine:latest
    command:
    - cat
    tty: true  
"""
}
  }
  stages {
    stage('Test') {
      steps {
        container('golang') {
          sh """
            ln -s `pwd` /go/src/sample-app
            cd /go/src/sample-app
            go test
          """
        }
      }
    }
    
    
    
    stage('Build and push image with Container Builder') {
      steps {
        container('alpine') {
            sh "docker build -t ${imageTag} ."
        }
      }
    }
    stage('Deploy Canary') {
      // Canary branch
      when { branch 'canary' }
      steps {
        container('kubectl') {

        } 
      }
    }
    stage('Deploy Production') {
      // Production branch
      when { branch 'master' }
      steps{
        container('kubectl') {
        
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
         
        }
      }     
    }
  }
}
}
