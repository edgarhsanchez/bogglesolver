FROM scratch

COPY main /main

COPY client/dist/client/* /public/
COPY hunspell/* /hunspell/

EXPOSE 8000

CMD ["/main"]

