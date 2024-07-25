FROM node:alpine as buildclient
# ARG GITHUB_SHA
# ARG SENTRY_AUTH_TOKEN

# ENV GITHUB_SHA=${GITHUB_SHA}
# ENV SENTRY_AUTH_TOKEN=${SENTRY_AUTH_TOKEN}

COPY client /build
RUN cd /build && npm i && npm run build



FROM golang:alpine as buildserver
COPY server /build
ENV GOPROXY=goproxy.cn,goproxy.io
RUN cd /build && go get && go build -o spc .


FROM alpine
EXPOSE 80
ENV ISDOCKER=true
VOLUME [ "/data" ]
COPY --from=buildserver /build/spc /app/spc
COPY --from=buildclient /build/dist /app/static

ENTRYPOINT [ "/app/spc" ]
