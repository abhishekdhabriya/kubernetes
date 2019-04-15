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

    stage ('Pull image') { 
      container('alpine') {
        
        
        sh 'uname -a'
        sh 'id'
        sh 'whoami'
        sh 'printenv'
        sh 'mount'
        sh 'ls -l /usr/bin/docker || true'
        sh 'ls -l /usr/local/bin/docker || true' 
        sh 'ls -l /var/run/docker.sock || true'
        sh 'which docker'
        
      }
    }
  }
}
