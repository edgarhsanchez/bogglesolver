FROM scratch

COPY main /main

COPY client/dist/client/* /public/
COPY hunspell/* /hunspell/

ENV PORT=80
ENV VERSION=.01

EXPOSE $PORT

CMD ["/main"]

