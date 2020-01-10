run:
	docker-compose up -d --force-recreate --build
	sleep 40
	# Fix pillar issue
	docker exec -it basic-salt-master salt '*' saltutil.refresh_pillar
	# Apply the hello state to the minion
	docker exec -it basic-salt-master salt '*' state.apply
	docker exec -it basic-salt-master salt '*' pillar.items

destroy:
	docker-compose down
