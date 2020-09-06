module "SimpleWebServer"

go 1.23

require (
	SimpleWebServer/utils v0.0.0
)

replace (
	SimpleWebServer/utils => ./utils
)