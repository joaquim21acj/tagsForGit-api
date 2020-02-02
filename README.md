# tagsForGit-api


## Mongodb

### Run mongo

If you have mongo already installed, you can just use your local configuration, localhost. But if you don't have it, you can run it with docker in this simples steps:

#### Update your system

Taking manjaro as an example, you should run:

sudo pacman -Syu # go/tagsForGit-api


## Mongodb

### Run mongo

If you have mongo already installed, you can just use your local configuration, localhost. But if you don't have it, you can run it with docker in this simples steps:

#### Installing docker

Taking manjaro as an example, you should run to update your system:
```sh
sudo pacman -Syu 
```
Then, run the command below to install, start and enable docker:
```sh
sudo pacman -S docker
sudo systemctl start docker
sudo systemctl enable docker
```

#### Installing mongo on docker

To install an image from mongo in to your docker, run this commands:
```sh
sudo docker pull tutum/mongodb
```
This next two commands are to create a container, the first is without password, the second is with password
```sh
sudo docker run -d -p 27017:27017 -p 28017:28017 -e AUTH=no tutum/mongodb
sudo docker run -d -p 27017:27017 -p 28017:28017 -e MONGODB_PASS="mypass" tutum/mongodb
```
The run the first command to find the id from mongo container, then use this id to execute the last command and it will be running and able to be accessed by the api.
```sh
sudo docker ps -a
sudo docker start id
```
