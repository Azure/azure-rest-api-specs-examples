package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7f70e351393addbc31d790a908c994c7c8644d9c/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-06-01-preview/examples/GetClusterInstanceView.json
func ExampleClustersClient_GetInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().GetInstanceView(ctx, "rg1", "clusterPool1", "cluster1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterInstanceViewResult = armhdinsightcontainers.ClusterInstanceViewResult{
	// 	Name: to.Ptr("default"),
	// 	Properties: &armhdinsightcontainers.ClusterInstanceViewResultProperties{
	// 		ServiceStatuses: []*armhdinsightcontainers.ServiceStatus{
	// 			{
	// 				Kind: to.Ptr("ZooKeeper"),
	// 				Message: to.Ptr(""),
	// 				Ready: to.Ptr("true"),
	// 		}},
	// 		Status: &armhdinsightcontainers.ClusterInstanceViewPropertiesStatus{
	// 			Message: to.Ptr(""),
	// 			Ready: to.Ptr("True"),
	// 			Reason: to.Ptr(""),
	// 		},
	// 	},
	// }
}
