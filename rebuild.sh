#!/bin/bash
sudo docker stop frogbot
sudo docker rm frogbot
sudo docker rmi frogbot:0.0.2
sudo docker build . -t frogbot:0.0.2 && sudo docker image prune -f
sudo docker create --restart=unless-stopped --name frogbot --env-file env frogbot:0.0.2
sudo docker start frogbot
