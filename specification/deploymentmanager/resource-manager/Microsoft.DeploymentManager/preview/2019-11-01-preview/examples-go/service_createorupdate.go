package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/service_createorupdate.json
func ExampleServicesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdeploymentmanager.NewServicesClient("caac1590-e859-444f-a9e0-62091c0f5929", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CreateOrUpdate(ctx,
		"myResourceGroup",
		"myTopology",
		"myService",
		armdeploymentmanager.ServiceResource{
			Location: to.Ptr("centralus"),
			Tags:     map[string]*string{},
			Properties: &armdeploymentmanager.ServiceResourceProperties{
				TargetLocation:       to.Ptr("centralus"),
				TargetSubscriptionID: to.Ptr("600c95c5-3ee5-44fe-b190-ca38a19adcd7"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
