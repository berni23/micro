FROM rabbitmq:3.11.5-alpine


# RUN mkdir /app



# RUN apt-get update \
#   && apt-get install -y --no-install-recommends \
#     ssh \
#     git \
#     m4 \
#     libgmp-dev \
#     opam \
#     wget \
#     ca-certificates \
#     rsync \
#     strace \
#     gcc \
#     rlwrap \
#     sudo


RUN useradd -ms /bin/bash  berni
RUN echo "berni:berni" | chpasswd
RUN adduser berni sudo

RUN chown berni /var/lib/rabbitmq
RUN chmod 777 /var/lib/rabbitmq -R






# RUN adduser -G root -u 1000 -h /home/berni berni --password 1234

# USER berni