FROM ubuntu:22.04

#RUN sudo apt-get update
RUN apt-get update  
RUN apt-get install -y ca-certificates
# RUN apt-get update \
#     && apt-get install -y \
#     ca-certificates 
ADD server /server
ADD entrypoint.sh /entrypoint.sh
WORKDIR /

EXPOSE 3000
ENTRYPOINT ["/entrypoint.sh"]

