package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/22ae5674fc98c32b29fb60791bd51a8fbd41b25f/specification/elastic/resource-manager/Microsoft.Elastic/stable/2025-06-01/examples/Monitors_Create.json
func ExampleMonitorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitorsClient().BeginCreate(ctx, "myResourceGroup", "myMonitor", &armelastic.MonitorsClientBeginCreateOptions{Body: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MonitorResource = armelastic.MonitorResource{
	// 	Name: to.Ptr("myMonitor"),
	// 	Type: to.Ptr("Microsoft.Elastic/monitors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
	// 	Kind: to.Ptr("elastic-hosted-deployment"),
	// 	Location: to.Ptr("West US 2"),
	// 	Properties: &armelastic.MonitorProperties{
	// 		ElasticProperties: &armelastic.Properties{
	// 			ElasticCloudDeployment: &armelastic.CloudDeployment{
	// 				Name: to.Ptr("deploymentname"),
	// 				AzureSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				DeploymentID: to.Ptr("deployment_id"),
	// 				ElasticsearchRegion: to.Ptr("azure-westus2"),
	// 				ElasticsearchServiceURL: to.Ptr("https://elasticsearchendpoint.com"),
	// 				KibanaServiceURL: to.Ptr("https://kibanaserviceurl.com"),
	// 				KibanaSsoURL: to.Ptr("https://kibanssourl.com"),
	// 			},
	// 			ElasticCloudUser: &armelastic.CloudUser{
	// 				ElasticCloudSsoDefaultURL: to.Ptr("https://examplessourl.com"),
	// 				EmailAddress: to.Ptr("alice@microsoft.com"),
	// 				ID: to.Ptr("myid123"),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armelastic.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armelastic.ResourceSKU{
	// 		Name: to.Ptr("free_Monthly"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}
