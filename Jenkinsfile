podTemplate(label: 'twistlock-example-builder', 
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
  node ('twistlock-example-builder') {

    stage ('Initiliaze Docker') { 
      container('alpine') {
        
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
