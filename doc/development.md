[Back](../README.md)
# Development Environment

# Linux
### Pre-requisites
install the following tools.
- [Docker](https://docs.docker.com/get-docker/)
- [helm](https://helm.sh/docs/intro/install/)
- [skaffold](https://skaffold.dev/docs/install/)
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

You might also want to install these awesome tools. Not required but they sure are epic! 
- [k9s](https://github.com/derailed/k9s)
- [kubectx](https://github.com/ahmetb/kubectx)
- [kubens](https://github.com/ahmetb/kubectx)

### Setup
There are two ways you can opt to setup the devwiki dev environment. You can either create cluster once using 
```bash 
./devcluster/setup.sh 
```
Then run skaffold to deploy the application when you want to work on it
```bash 
skaffold dev
```

Or you can run the `dev.sh` to bootstrap everything all at once. This will create a new cluster each time you run it.
and teardown the cluster when you're done. This is a great way to make sure you're always working with an updated & clean slate. 
```bash
./dev.sh
```

Whatever option you opt for you will also need to put an entry in your `/etc/hosts` file 
```
172.17.0.1 dev.devwiki.io
```


Now you should now be able to see devwiki running in your browser at [dev.devwiki.io](http://dev.devwiki.io) 🎉.

At this point you can start making changes to the code and skaffold will automatically sync your changes. The only notable 
exception is tailwindcss changes for server rendered pages. If you make changes that require new tailwinds classes then
you must regenerate the tailwindcss files. To do this run the following command from the `services/ui` directory.
```bash
./buildcss.sh
```
If you want you can add the `--watch-[dark|light]` flag to keep this command running and automatically regenerate the tailwinds css
file when you make changes that require additional tailwind classes.

# Windows 

### Step 1 
Install Linux. Don't worry about any partitions called "Windows" you won't need those anymore. Just overwrite them with Linux. 🤣

[Back](../README.md)