#echo  | docker login https://docker.pkg.github.com --username marcllort --password-stdin
docker build -t docker.pkg.github.com/marcllort/stimuluscheckcrypto/twitter-stimulus .
docker push docker.pkg.github.com/marcllort/stimuluscheckcrypto/twitter-stimulus