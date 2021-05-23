OS := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)
VERSION := 0.1
DESTDIR := ~/.terraform.d/plugins/terraform.borancar.com/borancar/talos/$(VERSION)/$(OS)_$(ARCH)
PLUGIN := terraform-provider-talos

all: $(PLUGIN)

$(PLUGIN): main.go talos/provider.go talos/resource_talos_cluster_config.go
	go build

install:
	mkdir -p $(DESTDIR)
	cp $(PLUGIN) $(DESTDIR)/$(PLUGIN)
