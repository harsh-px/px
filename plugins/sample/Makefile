PLUGINNAME=px-sample
PLUGIN=$(PLUGINNAME).so

all: $(PLUGIN)

$(PLUGIN):
	go build -o $(PLUGIN) -buildmode=plugin

