podTemplate(label: 'kube-by-example', 
  containers: [
    containerTemplate(
      name: 'jnlp',
      image: 'jenkinsci/jnlp-slave:3.10-1-alpine',
      args: '${computer.jnlpmac} ${computer.name}'
    ),
    containerTemplate(
      name: 'kubectl',
      image: 'gcr.io/cloud-builders/kubectl',
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
    
    stage ('Initiliaze Docker') { 
      container('jnlp') {
        def dockerHome = tool 'myDocker'
        env.PATH = "${dockerHome}/bin:${env.PATH}"
      }
    }

    stage ('Cloning Git') { 
      container('jnlp') {
        git 'https://github.com/anishnath/kubernetes.git'
        docker.build 'anishnath/hello'+ ":$BUILD_NUMBER"
      }
    }
  }
}
