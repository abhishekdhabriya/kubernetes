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
  node ('kube-by-example') {
    
    environment {
    registry = "anishnath/hello"
    registryCredential = 'docker-hub-creds'
    dockerImage = ''
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
    
    stage('Deploy Image') {
      container('alpine') {
          docker.withRegistry( '', 'docker-hub-creds' ) {
            dockerImage.push()
        }
      }
    }
  }
}
