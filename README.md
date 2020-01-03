**Task 0: Install a ubuntu 16.04 server 64-bit**

tools:  virtualBox 6.0

ubuntu image: http://releases.ubuntu.com/16.04/ubuntu-16.04.6-server-amd64.iso

refer to: https://www.wikihow.com/Install-Ubuntu-on-VirtualBox

proxy: http://www-proxy.ericsson.se:8080/

ubuntu： ip addr : 10.0.2.15

vb configure port：

         HostIP     hostPort           CuestIP GuestPort
         
        127.0.0.1   2222             10.0.2.15     22
        
        127.0.0.1   8080             10.0.2.15     80
        
        127.0.0.1   8081             10.0.2.15     8081
        
        127.0.0.1   8082             10.0.2.15     8082
        
        127.0.0.1   31080             10.0.2.15     31080
        
        127.0.0.1   31081             10.0.2.15     31081
        
**Task 1: Update system**

install git (version 2.24.1) in windows 

ssh eyinwag@127.0.0.1 -p 2222

update the system to the latest, upgrade the kernel to the 16.04  4.11.2

refer to: https://www.howtoforge.com/tutorial/how-to-upgrade-linux-kernel-in-ubuntu-1604-server/

Step 1 - Update Ubuntu Repository and Upgrade all Packages
ssh eyinwag@127.0.0.1 -p 2222
sudo -i
apt update
apt upgrade -y
reboot
ssh eyinwag@127.0.0.1 -p 2222
sudo apt list --upgradeable

Step 2 - Checking the Active Kernel Version
#uname -msr
Linux 4.4.0-170-generic x86_64

Step 3 - Install New Kernel Version
sudo mkdir -p ~/4.11.2
cd ~/4.11.2
path: find latest : https://kernel.ubuntu.com/~kernel-ppa/mainline/

set proxy：
export http_proxy=http://xxx

wget http://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.2/linux-headers-4.11.2-041102_4.11.2-041102.201705201036_all.deb
wget http://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.2/linux-headers-4.11.2-041102-generic_4.11.2-041102.201705201036_amd64.deb
wget http://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.2/linux-image-4.11.2-041102-generic_4.11.2-041102.201705201036_amd64.deb

dpkg -i *.deb
sudo update-grub
sudo reboot
ssh eyinwag@127.0.0.1 -p 2222
$ uname -msr
Linux 4.11.2-041102-generic x86_64

**Task 2: install gitlab-ce version in the host**

refer to: https://about.gitlab.com/install/#ubuntu
1. Install and configure the necessary dependencies
$sudo apt-get install -y curl openssh-server ca-certificates
2. Add the GitLab package repository and install the package
$curl -sS https://packages.gitlab.com/install/repositories/gitlab/gitlab-ee/script.deb.sh | sudo bash
$sudo EXTERNAL_URL="http://127.0.0.1" apt-get install gitlab-ee

**Task 3: create a demo group/project in gitlab**

http://127.0.0.1:8080/demo/go-web-hello-world
#apt install golang
#apt install git
#mkdir -p $HOME/go/src/hello
#export GOPATH=$HOME/go
#git clone  http://127.0.0.1/demo/go-web-hello-world.git
#cd go-web-hello-world
#vi main.go 
#go build
#ls
#./go-web-hello-world
verify: http://127.0.0.1:8081/
git add main.go
git status
git commit
git push origin master

**Task 4: build the app and expose ($ go run) the service to 8081 port**

$ curl http://127.0.0.1:8081
Hello World

**Task 5: install docker**

refer to:https://docs.docker.com/install/linux/docker-ce/ubuntu/
#sudo apt-get update
#sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
#curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
#sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
#sudo apt-get update
#sudo apt-cache madison docker-ce docker-ce-cli
#sudo apt install docker-ce=5:18.09.9~3-0~ubuntu-xenial docker-ce-cli=5:18.09.9~3-0~ubuntu-xenial
#docker version
Task 6: run the app in container
root@ubuntu:/home/emiosun/go-web-hello-world# docker build .
#docker build -t go-web-hello-world .
#docker images
#docker run -p 8083:8081 -it --rm --name run-web-hello-world go-web-hello-world
can access: http://127.0.0.1:8083

root@ubuntu:/home/emiosun# curl http://127.0.0.1:8083

Hello World

Task 7: push image to dockerhub
docker tag go-web-hello-world:latest amysunm/go-web-hello-world:v0.1
docker login
docker push amysunm/go-web-hello-world:v0.1
access : https://hub.docker.com/repository/docker/amysunm/go-web-hello-world

Task 12: generate token for dashboard login in task 11
refer to: https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md

Create Service Account

Create ClusterRoleBinding

Bearer Token

$kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | grep admin-user | awk '{print $1}')

