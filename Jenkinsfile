podTemplate(label: 'kube-by-example', 
  containers: [
    containerTemplate(
      name: 'jnlp',
      image: 'jenkinsci/jnlp-slave:3.10-1-alpine',
      args: '${computer.jnlpmac} ${computer.name}'
    ),
  ],
  volumes: [ 
    hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock'), 
  ]
)
{
  node ('kube-by-example') {
    
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
"""
}

    stage ('Initiliaze Docker') { 
      container('jnlp') {
        
        def dockerHome = tool 'myDocker'
        env.PATH = "${dockerHome}/bin:${env.PATH}"
      }
    }

    stage ('Cloning Git') { 
      container('alpine') {
        
        git 'https://github.com/anishnath/kubernetes.git'
        docker.build 'anishnath/hello'+ ":$BUILD_NUMBER"
      }
    }
      
      
  }
}
