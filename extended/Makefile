run:
	docker-compose up -d --force-recreate --build
	sleep 60
	# Fix pillar issue
	docker exec -it extended-salt-master salt '*' saltutil.refresh_pillar
	# Enroll minion with salt master
	docker exec -it extended-salt-master salt '*' state.apply
	docker exec -it extended-salt-master salt '*' pillar.items

destroy:
	# Remove named containers
	docker-compose down -v
