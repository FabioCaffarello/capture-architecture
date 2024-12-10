
lint:
	npx nx run-many --target=lint --all

check-all:
	npx nx run-many --target=test --all

tree:
	tree -I 'dist|node_modules'

cleanup:
	@max_retries=3; \
	attempt=0; \
	until [ $$attempt -ge $$max_retries ]; do \
		npx nx reset && break; \
		attempt=$$((attempt+1)); \
		echo "Retry $$attempt/$$max_retries..."; \
	done; \
	if [ $$attempt -ge $$max_retries ]; then \
		echo "Command failed after $$max_retries attempts."; \
	fi; \
	containers=$$(docker ps -q -a); \
	if [ -n "$$containers" ]; then \
		docker rm -f $$containers; \
	else \
		echo "No containers to remove"; \
	fi;