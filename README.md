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
### Many thanks to [Yevgeniy Brikman](https://github.com/brikis98)
For his excellent InfoQ talk [Automated Testing for Terraform, Docker, Packer, Kubernetes, and More](https://www.youtube.com/watch?v=xhHOW0EF5u8). This is a derivative of one of his examples: https://github.com/gruntwork-io/infrastructure-as-code-testing-talk.
