
SUBDIRS := $(shell find ./ -type d -not -path "*/.git*" -not -path "*pkg*" -not -path "*template*" -not -path "./" | sort )
.SILENT:
all: $(SUBDIRS)
	@echo $(SUBDIRS)

$(SUBDIRS):
	@echo "*************************$(basename $@)*************************"
	$(MAKE) -C $@ $(MAKECMDGOALS)

.PHONY: $(TOPTARGETS) $(SUBDIRS)