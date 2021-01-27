FROM instrumentisto/nmap:latest

WORKDIR /app
COPY sonar .

ENTRYPOINT ["./sonar"]