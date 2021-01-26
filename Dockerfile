FROM instrumentisto/nmap:latest

WORKDIR /app
COPY mri .

ENTRYPOINT ["./mri"]