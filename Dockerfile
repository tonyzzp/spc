FROM node:alpine as buildclient
# ARG GITHUB_SHA
# ARG SENTRY_AUTH_TOKEN

# ENV GITHUB_SHA=${GITHUB_SHA}
# ENV SENTRY_AUTH_TOKEN=${SENTRY_AUTH_TOKEN}

COPY client /build
RUN cd /build && npm i && npm run build



FROM golang:alpine as buildserver
COPY server /build
RUN cd /build && go get && go build -o server .


FROM alpine
EXPOSE 80
ENV ISDOCKER=true
VOLUME [ "/data" ]
COPY --from=buildserver /build/server /app/server
COPY --from=buildclient /build/dist /app/static

CMD ["/app/server"]