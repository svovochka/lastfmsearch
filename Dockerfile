FROM library/golang:1.14.13
RUN mkdir /app
WORKDIR /app
COPY ./lastfmsearch ./lastfmsearch
CMD ["/app/lastfmsearch"]