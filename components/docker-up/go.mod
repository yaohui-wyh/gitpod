module github.com/gitpod-io/gitpod/docker-up

go 1.18

require (
	github.com/opencontainers/runtime-spec v1.0.2
	github.com/rootless-containers/rootlesskit v1.0.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/vishvananda/netlink v1.1.0
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

require github.com/vishvananda/netns v0.0.0-20191106174202-0a2b9b5464df // indirect
