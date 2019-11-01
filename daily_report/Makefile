SHELL:=/bin/bash
DATE:=$$(date +%Y%m%d)
SCHEDULED:=true
LATEST:=$$(ls -t | grep .md | head -n 1)
YESTERDAY:=$$(date +%Y%m%d -d "1 day ago")

define _rmsuffix
endef


create:
	$(eval FILENAME := $$(DATE).md)
	touch $(FILENAME)
	echo "# $$(date -d $(DATE) +%Y/%m/%d)" >> $(FILENAME)
	echo "## 今日のトレーニングメニュー" >> $(FILENAME)
	if [ $(SCHEDULED) = true ]; then\
		tools/bin/schedule schedule.dat >> $(FILENAME);\
	fi
	echo "## 所感" >> $(FILENAME)
	code $(FILENAME)

create/yesterday:
	$(MAKE) create DATE=$(YESTERDAY)

save:
	$(eval FILENAME := $$(DATE).md)
	git add $(FILENAME)
	git commit -m "$(DATE)"
	git push -u

save/yesterday:
	$(MAKE) save DATE=$(YESTERDAY)

open/latest:
	code $(LATEST)

remove/latest:
	rm $(LATEST)