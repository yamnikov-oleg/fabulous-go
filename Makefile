SHELL=/bin/bash
WD=./tmp
RAIDER_BIN=$(WD)/raider
RAIDER_DIR=./raider
EXHIBITOR_BIN=$(WD)/exhibitor
EXHIBITOR_DIR=./exhibitor
TEMPLATE=TEMPLATE.md
README=README.md

all: raid exhibit

raid: prepare_raider run_raider cleanup_raider

prepare_raider:
	go build -o $(RAIDER_BIN) $(RAIDER_DIR)

run_raider:
	cd $(WD) && ../$(RAIDER_DIR)/run_authorized.sh

cleanup_raider:
	-rm $(RAIDER_BIN)

exhibit: prepare_exhibitor run_exhibitor cleanup_exhibitor

prepare_exhibitor:
	go build -o $(EXHIBITOR_BIN) $(EXHIBITOR_DIR)
	cp $(EXHIBITOR_DIR)/$(TEMPLATE) $(WD)/$(TEMPLATE)

run_exhibitor:
	cd $(WD) && ../$(EXHIBITOR_BIN)
	cp $(WD)/$(README) ./$(README)

cleanup_exhibitor:
	-rm $(EXHIBITOR_BIN)
	-rm $(WD)/$(README)
	-rm $(WD)/$(TEMPLATE)

cleanup: cleanup_raider cleanup_exhibitor
