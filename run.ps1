docker rm -f vfs-app
docker run --name vfs-app --link selenium-hub vfs