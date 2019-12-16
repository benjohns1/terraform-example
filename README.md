# Terraform Example
Need to have:
 - Go 1.13
 - Terraform 0.12
 - AWS account and local credential file
### Build
In `app/build`
```shell script
make build
```
(or `build.cmd` on Windows)
### Deploy
In `app/deploy`
```shell script
terraform init
terraform apply
```
Go to the output URL, manually add `?name=YourName` to the URL.
### Tear Down
In `app/deploy`
```shell script
terraform destroy
```