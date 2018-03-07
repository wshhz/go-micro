build:
    docker build -t vessel-service .

run:
    docker run -p 50052:50051 -e MICRO_SERVER_ADDRES=:50051 -e MICRO_REGISTRY=mdns -d vessel-service