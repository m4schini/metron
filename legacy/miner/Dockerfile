# miner docker file

# building go executable in own stage
FROM golang:1.17.5 AS build

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o miner .

# using headless-shell as base image provides the chrome instance
FROM chromedp/headless-shell:latest

# installing curl, needed for healtcheck
RUN apt update
RUN apt install dumb-init curl -y

#TODO nordvpn installation not working: 'We could not disable IPv6. For more information, please check activity log.'
# installing nordvpn in container
# ARG VERSION=3.7.4
# RUN apt update && \
#        apt install -y wget && \
#        wget -O /tmp/nordrepo.deb https://repo.nordvpn.com/deb/nordvpn/debian/pool/main/nordvpn-release_1.0.0_all.deb && \
#        apt install -y /tmp/nordrepo.deb && \
#        apt update && \
#        apt install -y nordvpn=$VERSION && \
#        apt remove -y wget nordvpn-release
# ENTRYPOINT ["/usr/sbin/nordvpnd", "&"]

# disabling ipv6, required by nordvpn
# RUN echo "net.ipv6.conf.all.disable_ipv6 = 1" >> /etc/sysctl.conf
# RUN echo "net.ipv6.conf.default.disable_ipv6 = 1" >> /etc/sysctl.conf
# RUN echo "net.ipv6.conf.lo.disable_ipv6 = 1" >> /etc/sysctl.conf
# RUN echo "net.ipv6.conf.tun0.disable_ipv6 = 1" >> /etc/sysctl.conf

# copying executable from build stage
RUN mkdir /app
COPY --from=build /app/miner /app/miner

# exposing required ports
EXPOSE 50051
EXPOSE 8080

# running init, required by headless shell to avoid zombie processes
ENTRYPOINT ["dumb-init", "--"]

# using local mode
CMD ["/app/miner"]