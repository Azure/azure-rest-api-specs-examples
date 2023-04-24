package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/Monitors_ListByResourceGroup.json
func ExampleMonitorsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListByResourceGroupPager("myResourceGroup", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.MonitorResourceListResponse = armelastic.MonitorResourceListResponse{
		// 	Value: []*armelastic.MonitorResource{
		// 		{
		// 			Name: to.Ptr("myMonitor"),
		// 			Type: to.Ptr("Microsoft.Elastic/monitors"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
		// 			Location: to.Ptr("West US 2"),
		// 			Properties: &armelastic.MonitorProperties{
		// 				ElasticProperties: &armelastic.Properties{
		// 					ElasticCloudDeployment: &armelastic.CloudDeployment{
		// 						Name: to.Ptr("deploymentname"),
		// 						AzureSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						DeploymentID: to.Ptr("deployment_id"),
		// 						ElasticsearchRegion: to.Ptr("azure-westus2"),
		// 						ElasticsearchServiceURL: to.Ptr("https://elasticsearchendpoint.com"),
		// 						KibanaServiceURL: to.Ptr("https://kibanaserviceurl.com"),
		// 						KibanaSsoURL: to.Ptr("https://kibanssourl.com"),
		// 					},
		// 					ElasticCloudUser: &armelastic.CloudUser{
		// 						ElasticCloudSsoDefaultURL: to.Ptr("https://examplessourl.com"),
		// 						EmailAddress: to.Ptr("alice@microsoft.com"),
		// 						ID: to.Ptr("myid123"),
		// 					},
		// 				},
		// 				LiftrResourceCategory: to.Ptr(armelastic.LiftrResourceCategoriesMonitorLogs),
		// 				LiftrResourcePreference: to.Ptr[int32](0),
		// 				MonitoringStatus: to.Ptr(armelastic.MonitoringStatusEnabled),
		// 				ProvisioningState: to.Ptr(armelastic.ProvisioningStateSucceeded),
		// 			},
		// 			Tags: map[string]*string{
		// 				"Environment": to.Ptr("Dev"),
		// 			},
		// 	}},
		// }
	}
}
