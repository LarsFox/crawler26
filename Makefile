include .env

default:
	@printf "Start with 'make init' or 'cat Makefile' to see what you can do."

init:
	@cp .env.sample .env
	@printf "Fill the .env file then 'make auth'.\n"

auth:
	@printf "To get access token, follow https://oauth.vk.com/authorize?client_id=${VK_API_APP_ID}&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=offline&response_type=token&v=5.52\n"
	@printf "Full list of permissions: https://vk.com/dev/permissions\n"

gorun:
	@go run go/cmd/main.go --vk_api_token $(VK_API_ACCESS_TOKEN)

j:
	jupyter-notebook

py_get_pikabu:
	@python3 python/pikabu_get.py

py_vk:
	@python3 python/vk.py --token $(VK_API_ACCESS_TOKEN)
