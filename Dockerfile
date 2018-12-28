FROM iron/base

COPY vitclient /opt/service/

WORKDIR /opt/service

RUN chmod +x /opt/service/vitclient

EXPOSE 8080:8080

CMD /opt/service/vitclient