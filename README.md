# gaufre

gaufre is an answering machine bot for IRC that forwards messages to pushbullet when your status is away

#### run on docker

```
docker build -t gaufre_img_1 .
docker run -d --name gaufre_cont_1 --net=host gaufre_img_1
```
