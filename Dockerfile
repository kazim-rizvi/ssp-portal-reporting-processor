FROM golang:1.21.2-alpine3.18

ARG PROFILE=none

WORKDIR /ssp-portal-reporting-processor

COPY . .

RUN go build -o reporting-processor-app ./cmd

ENV PROJECT_VERSION="1.0.0"

ENV PROFILE=${PROFILE}

ENV FILE_DIR /var/data

RUN mkdir -p $FILE_DIR

RUN chown -R 1000:1000 $FILE_DIR

RUN chmod +x reporting-processor-app

USER 1000:1000

CMD ["./reporting-processor-app"]
