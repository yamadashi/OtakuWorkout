DATE:=`date +%Y%m%d`
FILENAME:=$(DATE).md

create:
	touch $(FILENAME)
	echo "# `date +%Y/%m/%d`" >> $(FILENAME)
	echo "## 今日のトレーニングメニュー" >> $(FILENAME)
	echo "## 所感" >> $(FILENAME)
	code $(FILENAME)

create/yesterday:
	$(eval DATE := `date +%Y%m%d --date "1 day ago"`)
	$(eval FILENAME := $(DATE).md)
	touch $(FILENAME)
	echo "# `date +%Y/%m/%d --date "1 day ago"`" >> $(FILENAME)
	echo "## 今日のトレーニングメニュー" >> $(FILENAME)
	echo "## 所感" >> $(FILENAME)
	code $(FILENAME)

save:
	git add $(FILENAME)
	git commit -m "$(DATE)"
	git push -u

save/yesterday:
	$(eval DATE := `date +%Y%m%d --date "1 day ago"`)
	$(eval FILENAME := $(DATE).md)
	git add $(FILENAME)
	git commit -m "$(DATE)"
	git push -u