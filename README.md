# Deploy Go app to AWS lambda

Our application needs to delete existing ClusterNodeGroups with Specific Tags at specific times


## Step 1: Write the Go code

download eks library from GitHub:
```
go get github.com/aws/aws-lambda-go/eks
```
The code of the main function
```
func main() {

	var ClusList = ListClusters()
	for _, e := range ClusList {
		var DesClus = DescribeCluster(*e)
		for _, i := range DesClus {
			var Tag = *i
			if Tag == "true" {
				var ClusName = *e
				var no = ListNodeGroup(ClusName, 2)
				for _, a := range no {
					fmt.Println(*a)
					var NodeName = *a
					fmt.Println(ClusName)
					fmt.Println(NodeName)
					var del = DeleteNodeGroupe(ClusName, NodeName)
					fmt.Println(*del)

				}
			}

		}

	}

}
```
The rest of the code is in the main.go file

Compile your executable file:
```
$ GOOS=linux go build main.go
```

# Step 2: Create a Go Image and Deploy to ECR

Use a text editor to create a Dockerfile in your project directory. The following example Dockerfile uses the AWS provided.al2 base image.

```
FROM public.ecr.aws/lambda/provided:al2 as build
# install compiler
RUN yum install -y golang
RUN go env -w GOPROXY=direct
# cache dependencies
ADD go.mod go.sum ./
RUN go mod download
# build
ADD . .
RUN go build -o /main
# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /main /main
ENTRYPOINT [ "/main" ]           
```
Build your Docker image with the docker build command. Enter a name for the image. 
```
docker build -t image_name . 
```
Authenticate the Docker CLI to your Amazon ECR registry.
```
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 123456789012.dkr.ecr.us-east-1.amazonaws.com    
```
Tag your image to match your repository name, and deploy the image to Amazon ECR using the docker push command.
```
docker tag  hello-world:latest 123456789012.dkr.ecr.us-east-1.amazonaws.com/hello-world:latest
docker push 123456789012.dkr.ecr.us-east-1.amazonaws.com/hello-world:latest           
```
Now that your container image resides in the Amazon ECR container registry, you can create the Lambda function and deploy the image.

# Step 3: Create the lambda function

Open the Functions page of the Lambda console.

Choose Create function.

Choose the Container image option.

Under Basic information, do the following:

For Function name, enter the function name.

For Container image URI, provide a container image that is compatible with the instruction set architecture that you want for your function code.

You can enter the Amazon ECR image URI or browse for the Amazon ECR image.

Enter the Amazon ECR image URI.

Or, to browse an Amazon ECR repository for the image, choose Browse images. Select the Amazon ECR repository from the dropdown list, and then select the image.


# Etape 4:  Create a Rule to launch Lambda

Open the Amazon EventBridge console at https://console.aws.amazon.com/events/.

In the navigation pane, choose Rules.

Choose Create rule.

Enter a name and description for the rule.

A rule can't have the same name as another rule in the same Region and on the same event bus.

For Event bus, choose the event bus that you want to associate with this rule. If you want this rule to match events that come from your account, select AWS default event bus. When an AWS service in your account emits an event, it always goes to your account’s default event bus.

For Rule type, choose Schedule.

Choose Next.

For Schedule pattern, choose A schedule that runs at a regular rate, such as every 10 minutes. and enter 5 and choose Minutes from the drop-down list.

Choose Next.

For Target types, choose AWS service.

For Select a target, choose Lambda function from the drop-down list.

For Function, select the Lambda function that you created in the Step 3: Create a Lambda function section. In this example, select LogScheduledEvent.

Choose Next.

Choose Next.

Review the details of the rule and choose Create rule.

