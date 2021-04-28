build: 
	ansible-playbook -v -e build_images=true -e push_images=false -e make_latest=false build/build-playbook.yml

test:
	./test/integration_tests.sh

release: 
	ansible-playbook -v -e build_images=true -e push_images=true -e make_latest=false build/build-playbook.yml

latest:
	ansible-playbook -v -e build_images=true -e push_images=true -e make_latest=true build/build-playbook.yml

.PHONY: build test
