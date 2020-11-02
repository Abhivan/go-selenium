docker-selenium-up:
	docker run -d -p 4444:4444 --name selenium-hub selenium/hub
	docker run -d --link selenium-hub:hub selenium/node-chrome
	docker run -d --link selenium-hub:hub selenium/node-firefox
build:
	docker build -t vfs .
remove: 
	docker rm -f vfs-app
run:
	docker run --name vfs-app --link selenium-hub vfs go run main.go