#!/usr/bin/env bash
docker-compose up -d
INSTANCE_ID=$(docker-compose exec moto_server \
  aws --profile=dummy_profile --region=us-east-1 --endpoint=http://127.0.0.1:5000 \
    ec2 run-instances \
      --image-id=ami-1a2b3c4d \
      --count=1 \
      --instance-type=c3.large \
      --key-name=MyKeyPair \
      --security-groups=MySecurityGroup \
      --query=Instances[].InstanceId --output=text)
docker-compose exec moto_server \
  aws --profile=dummy_profile --region=us-east-1 --endpoint=http://127.0.0.1:5000 \
    ec2 create-image \
      --instance-id=$(echo ${INSTANCE_ID} | tr -d \\r) \
      --name=test-image --output=text