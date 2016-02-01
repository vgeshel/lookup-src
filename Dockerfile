FROM scratch
MAINTAINER vgeshel@gmail.com
ADD lookup-srv /
ENTRYPOINT ["/lookup-srv"]
