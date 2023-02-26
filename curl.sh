curl 'https://registry.hub.docker.com/v2/library/ubuntu/manifests/latest' \
    -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
    -H "Authorization: Bearer $TOKEN"
