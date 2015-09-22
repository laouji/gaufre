FROM golang

ADD . $GOPATH/src/github.com/laouji/gaufre

WORKDIR $GOPATH/src/github.com/laouji/gaufre
RUN go get github.com/thoj/go-ircevent
RUN go get gopkg.in/yaml.v2
RUN go install .

CMD ["gaufre"]
