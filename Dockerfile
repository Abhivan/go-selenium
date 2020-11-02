FROM golang

RUN go get github.com/tebeka/selenium

RUN apt-get update && apt-get install nano sendmail net-tools -y

COPY . /app

WORKDIR /app

RUN chmod +x dockerscriptetchosts.sh && ./dockerscriptetchosts.sh

CMD /etc/init.d/sendmail start && go run main.go