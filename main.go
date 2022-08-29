package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

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
func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To delete a cluster
// This example command deletes a cluster named `devel` in your default region.
func DeleteNodeGroupe(cluster_name string, node_name string) *string {
	svc := eks.New(session.New())
	input := &eks.DeleteNodegroupInput{
		ClusterName:   aws.String(cluster_name),
		NodegroupName: aws.String(node_name),
	}

	result, err := svc.DeleteNodegroup(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case eks.ErrCodeResourceInUseException:
				fmt.Println(eks.ErrCodeResourceInUseException, aerr.Error())
			case eks.ErrCodeResourceNotFoundException:
				fmt.Println(eks.ErrCodeResourceNotFoundException, aerr.Error())
			case eks.ErrCodeClientException:
				fmt.Println(eks.ErrCodeClientException, aerr.Error())
			case eks.ErrCodeServerException:
				fmt.Println(eks.ErrCodeServerException, aerr.Error())
			case eks.ErrCodeServiceUnavailableException:
				fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		//return
	}

	//fmt.Println(result)
	return result.Nodegroup.Status
}

// To describe a cluster
// This example command provides a description of the specified cluster in your default
// region.

func DescribeCluster(cluster_name string) map[string]*string {
	svc := eks.New(session.New())

	input := &eks.DescribeClusterInput{
		Name: aws.String(cluster_name),
	}

	result, err := svc.DescribeCluster(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case eks.ErrCodeResourceNotFoundException:
				fmt.Println(eks.ErrCodeResourceNotFoundException, aerr.Error())
			case eks.ErrCodeClientException:
				fmt.Println(eks.ErrCodeClientException, aerr.Error())
			case eks.ErrCodeServerException:
				fmt.Println(eks.ErrCodeServerException, aerr.Error())
			case eks.ErrCodeServiceUnavailableException:
				fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		//return
	}

	return result.Cluster.Tags
}

// To list your available clusters
// This example command lists all of your available clusters in your default region.
func ListClusters() []*string {
	svc := eks.New(session.New())
	input := &eks.ListClustersInput{}

	result, err := svc.ListClusters(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case eks.ErrCodeInvalidParameterException:
				fmt.Println(eks.ErrCodeInvalidParameterException, aerr.Error())
			case eks.ErrCodeClientException:
				fmt.Println(eks.ErrCodeClientException, aerr.Error())
			case eks.ErrCodeServerException:
				fmt.Println(eks.ErrCodeServerException, aerr.Error())
			case eks.ErrCodeServiceUnavailableException:
				fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		//return
	}

	/*for _, e := range test {
		fmt.Println(*e)
	}*/
	return result.Clusters

	///fmt.Println(reflect.ValueOf(result).Kind())

}

func ListNodeGroup(cluster_name string, val int64) []*string {
	svc := eks.New(session.New())
	input := &eks.ListNodegroupsInput{
		ClusterName: aws.String(cluster_name),
		MaxResults:  aws.Int64(val),
	}
	result, err := svc.ListNodegroups(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case eks.ErrCodeInvalidParameterException:
				fmt.Println(eks.ErrCodeInvalidParameterException, aerr.Error())
			case eks.ErrCodeClientException:
				fmt.Println(eks.ErrCodeClientException, aerr.Error())
			case eks.ErrCodeServerException:
				fmt.Println(eks.ErrCodeServerException, aerr.Error())
			case eks.ErrCodeServiceUnavailableException:
				fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		//return
	}

	/*for _, e := range test {
		fmt.Println(*e)
	}*/
	//return result.Nodegroups
	//fmt.Println(result,err)
	//fmt.Printf("%s\n", result)
	return result.Nodegroups

}

type JsonType struct {
	Array []string
}

// To list tags for a cluster
// This example lists all of the tags for the `beta` cluster.
func ExampleEKS_ListTagsForResource_shared00() {
	svc := eks.New(session.New())
	input := &eks.ListTagsForResourceInput{
		ResourceArn: aws.String("arn:aws:eks:us-west-2:012345678910:cluster/beta"),
	}

	result, err := svc.ListTagsForResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case eks.ErrCodeBadRequestException:
				fmt.Println(eks.ErrCodeBadRequestException, aerr.Error())
			case eks.ErrCodeNotFoundException:
				fmt.Println(eks.ErrCodeNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
