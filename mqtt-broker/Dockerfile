FROM rabbitmq:3.8-management
RUN rabbitmq-plugins enable --offline rabbitmq_mqtt
RUN apt-get -y update && apt-get install netcat -y
