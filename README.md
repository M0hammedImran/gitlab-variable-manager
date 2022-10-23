# gitlab-variable-manager

CLI to Query for all the CI/CD variables from GitLab.

### Usage
```bash
git clone https://github.com/M0hammedImran/gitlab-variable-manager.git
cd gitlab-variable-manager
go build main.go  
```
```
./main <gitlabHost> <projectId> <privateToken>
```
```bash 
./main https://gitlab.com 1234 pasjdk_sdadd
```

- gitlabHost can even be a self hosted instance like `https://gitlab.example.dev`.

### TODO
- [ ] Query based on environment.
- [ ] Query for a set of varaibles.
- [ ] Query for one variable across different environments.
- [ ] Add one variable.
- [ ] Add many variables.
- [ ] Update a variable.
- [ ] Delate a variable.
