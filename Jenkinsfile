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
  - name: alpine
    image: twistian/alpine:latest
    command:
    - cat
    tty: true
  - name: kubectl
    image: gcr.io/cloud-builders/kubectl
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
            ln -s `pwd` /go/src/hello-app
            cd /go/src/hello-app
            go test
          """
        }
      }
    }
    stage('Build and push image with Container Builder') {
      steps {
        container('alpine') {
          docker.build "anishnath/hello" + ":$BUILD_NUMBER"
        }
      }
    }
  }
}
