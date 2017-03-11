# Swift Cleanr

A simple program used to clean an openstack swift storage container.


## Build


Run :

```bash
make
```

And that's all.


## Use

It required a yaml config like this :

```yaml
openstack:
  username: "AAAAAAAA"
  api_key: "BBBBBBBBBBB"
  auth_url: "https://auth.cloud.ovh.net/v2.0/"
  authtenant_name: "CCCCCCCCC"
  region: "SBG1"
```

Then you case use it in this way :

```bash
swift_cleanr -f config.yml -c container_name
```

By default it look at a config.yml in the same folder and cleanup the screencapture-staging containers.
