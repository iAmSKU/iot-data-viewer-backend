rm -rf iot-data-viewer-backend.tar
docker rmi iot-data-viewer-backend:0.0.1

docker build -f Dockerfile . -t iot-data-viewer-backend:0.0.1

docker save -o iot-data-viewer-backend.tar iot-data-viewer-backend:0.0.1
chmod 755 iot-data-viewer-backend.tar
