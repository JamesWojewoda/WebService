Go Web Service

This Web Service is set up to post strings and their sha256 encrypted values into a map. The map's data is ephemeral in nature
and will last as long as the docker container runs. Also, the web service is set up to get string values based on sha256 values
received via the respective endpoint. If nothing is returned a 404 will be served.
