# Sample API in Go hosted in AWS ECS

![Build and Deploy workflow status](https://github.com/orlowskilp/aws-ecs-api-go/workflows/Build%20and%20Deploy/badge.svg)
![Quick build workflow status](https://github.com/orlowskilp/aws-ecs-api-go/workflows/Quick%20check/badge.svg)
![Container build workflow status](https://github.com/orlowskilp/aws-ecs-api-go/workflows/Check%20Docker/badge.svg)
![CFN template validation workflow status](https://github.com/orlowskilp/aws-ecs-api-go/workflows/Check%20IaC/badge.svg)

This repository contains code and CI/CD pipeline demonstrating how to deploy a simple
API implemented in Go to AWS ECR.

## Development environment for [VSCode](https://code.visualstudio.com/)

My team and I love using [Microsoft Visual Studio Code](https://code.visualstudio.com/). We also, typically, develop on different platforms than our target platforms. VSCode
allows us to encapsulate the development environment, by leveraging the power of containers.

```
.devcontainer
├── Dockerfile
└── devcontainer.json
```

The `.devcontainer` directory contains the `Dockerfile` and development container 
orchestration file.  
Install `ms-vscode-remote.remote-containers` in your VSCode IDE and then open the 
repository directory. VSCode will offer to reopen the directory in container.
This gives you a development environment with all the dependencies pre-installed.

## API server code

The simple API server is implemented with
[gin-gonic/gin](https://github.com/gin-gonic/gin) library. It exposes 2 GET methods:
* `/kernel` - displays kernel information
* `/hostname` - displays hostname

There are 2 packages:
* `sys` - implements logic talking to the operating system
* `router` - implements a trivial HTTP server

By default the server listens on port `8080`. Port can be specified as a runtime
parameter.

```
go.mod
go.sum
main.go
pkg
├── router
│   ├── router.go
│   └── router_test.go
└── sys
    ├── sys.go
    └── sys_test.go
```

## Code packaged in Docker container

The `Dockerfile` assembles an image using multistage build, to keep the API image lean.
The resulting image which is uploaded to ECR is around 20MB in size.

The following parameters can be set during build time (with `--build-arg`):
* `VERSION` - server version
* `NAME` - server name
* `PORT` - port the server listens on

## AWS CloudFormation templates for infrastructure hosted in AWS

The `cfn` directory contains CloudFormation (CFN) templates to orchestrate all
of the necessary AWS resources. The templates in the `cicd` directory build stacks
necessary to get the CI/CD pipelines going.

```
cfn
├── cicd
│   ├── artifact-bucket.yml
│   ├── principal.yml
│   └── repository.yml
└── ecs
    └── cluster.yml
```

There are templates for four stacks:
* S3 bucket for tarballs
* ECR repository for container images
* IAM managed policy and user to authenticate as
* ECS resources (cluster, service and task definition) and required IAM role and VPC resources

There are four templates to make infrastructure management easier.

**NOTE**: Once there are objects in S3 bucket or ECR repository you won't be able
to delete the bucket or repository, until the objects are removed

## CI/CD workflows

The pipeline consists of multiple workflows:
* _Quick check_ - builds the server and runs unit tests on the code
* _Check container build_ - builds code and packs it into container
* _Check CloudFormation templates_ - validates if CFN templates are well-formed
* _Build and Deploy_ - builds, tests and packages the code and then packs it in
container and deploys it.

```
.github
├── CODEOWNERS
└── workflows
    ├── build.yml
    ├── check.yml
    ├── infra-check.yml
    └── pack-check.yml
```

The first three workflows, given their simple nature, are triggered on every push.

_Build and Deploy_ is triggered only on pushes to `master`, `staging` and `cicd-*`
branches as well as on Pull Requests on those branches.
_Build and Deploy_ runs the following jobs:
* Archive code to tarball
* Build code
* Run unit tests
* Deploy tarball to S3
* Pack application into container
* Push container to ECR repository
* Update container version in ECS task definition and update ECS service

**NOTE:** _Build and Deploy_ workflow requires AWS resources.

After successful execution, the API code runs in AWS ECS and can be accessed over
the browser. 

**NOTE:** To keep things simple and cost-effective, no load balancer is placed
in front of the ECS service, therefore you'll need to go to your AWS account and
fetch public IP addresses of the tasks to access the API in the browser.

### Required AWS resources

CI/CD pipeline will require the following AWS resources to pass:
* IAM policy and user - to access AWS resources
* S3 Bucket - to store artifact tarballs
* ECR Repository - to store API server containers
* ECS Cluster - to deploy service to
* ECS Service - to aggregate tasks running in containers
* ECS Task Execution Role - to allow task perform ECS specific actions
* ECS Task Definition - to describe configuration and container image to run

## Environment variables and secrets

You will need to add GitHub sercrets before the pipeline can talk to the AWS resources.
You might also want to customize the environment variabls.

### Required GitHub secrets

The pipeline requires the following sensitive data be stored in GitHub secrets:
* `AWS_ACCESS_KEY_ID` - AWS access key ID required to authenticate as IAM user
* `AWS_SECRET_ACCESS_KEY` - AWS secret access key required to authenticate as IAM user
* `AWS_ACCOUNT_NO` - AWS account number
* `ORG_DOMAIN` - DNS domain name of your organization
* `S3_BUCKET_NAME` - S3 bucket URI

At the very least, the first two secrets must not be compromised!

### Required environment variables

You may want to update the following environment variables in workflow files:

The default region is set to Singapore (where I'm based). If you built the
CloudFormation stacks in a different region, you'll need to update the
`AWS_DEFAULT_REGION` accordingly.

Pass the following values to the corresponding parameters in the CFN template (or 
update them in the workflow document, if you want different values):
* `SERVICE_NAME: sampleapi`
* `SERVICE_PORT: 8080`
* `AWS_ECS_TASK_NAME: sampleapi-ecs-task`
* `AWS_ECS_SERVICE_NAME: sampleapi-ecs-service`
* `AWS_ECS_CLUSTER_NAME: sampleapi-ecs-cluster`

## Getting everything running

You'll need to perform the following steps for the CI/CD pipeline to pass

**NOTE**: Take a note of parameters you pass to CFN templates, because they will
need to match between templates

1. Create a stack from the `cfn/cicd/artifacts-bucket.yml` template
2. Create a stack from the `cfn/cicd/repository.yml` template
3. Create a stack from the `cfn/cicd/principal.yml` template
4. Generate access keys for the created user
5. Set `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` GitHub repository secrets
6. Set `AWS_ACCOUNT_NO`, `ORG_DOMAIN`, `S3_BUCKET_NAME` GitHub repository secrets
7. Run the _Build and Deploy_ workflow (e.g. by pushing empty `staging` branch). It
will fail, but before that it will populate the ECR repository, which is required
to create an ECS service.
8. Create a stack from the `cfn/ecs/cluster.yml` template