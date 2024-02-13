FROM alpine:latest

RUN mkdir /app

COPY mailApp /app

COPY teamplete /teamplete
COPY teamplete /teamplete

CMD [ "/app/mailApp" ]