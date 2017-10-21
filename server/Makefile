.DEFAULT_GOAL := help

docker_server: docker_build docker_up

docker_clean: docker_stop docker_rm

help:
	@echo docker_build:	dockerのコンテナをビルドします
	@echo docker_up:	dockerのコンテナを立ち上げます
	@echo docker_stop:	dockerのコンテナを止めます
	@echo docker_rm:	dockerのコンテナを削除します
	@echo docker_ssh:	dockerのコンテナにsshでアクセスします
	@echo docker_restart: dockerのコンテナを再起動します

docker_build:
	docker-compose build

docker_up:
	docker-compose up

docker_stop:
	docker-compose stop

docker_restart:
	docker-compose restart

docker_rm:
	docker-compose rm

docker_ssh:
	docker exec -it plus-api /bin/bash

direnv:
	direnv allow

SERVER := "./go/src/app"

install:
	$(MAKE) -C $(SERVER) submodule
	$(MAKE) -C $(SERVER) depend
	$(MAKE) -C $(SERVER) gen

mac-install:
	which direnv || brew install direnv
	@echo 'zshならzshrcに eval "$$(direnv hook zsh)" の記述を追加してください'
	@echo 'bashならbashrcに　eval "$$(direnv hook bash)"  の記述を追加してください'
	@echo 参考URL:https://qiita.com/ngyuki/items/fda1bbf29384bef7a805

win-install:
	git clone http://github.com/zimbatm/direnv
	cd direnv
	sudo make install
	@echo bashrcに　eval "$(direnv hook bash)"  の記述を追加してください
	@echo 参考URL:https://qiita.com/ngyuki/items/fda1bbf29384bef7a805