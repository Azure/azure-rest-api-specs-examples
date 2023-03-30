package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/services_list.json
func ExampleServicesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().List(ctx, "myResourceGroup", "myTopology", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResourceArray = []*armdeploymentmanager.ServiceResource{
	// 	{
	// 		Name: to.Ptr("Service East"),
	// 		Type: to.Ptr("Microsoft.DeploymentManager/serviceTopologies/services"),
	// 		Location: to.Ptr("centralus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armdeploymentmanager.ServiceResourceProperties{
	// 			TargetLocation: to.Ptr("eastus"),
	// 			TargetSubscriptionID: to.Ptr("600c95c5-3ee5-44fe-b190-ca38a19adcd7"),
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("Service West"),
	// 		Type: to.Ptr("Microsoft.DeploymentManager/serviceTopologies/services"),
	// 		Location: to.Ptr("centralus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armdeploymentmanager.ServiceResourceProperties{
	// 			TargetLocation: to.Ptr("westus"),
	// 			TargetSubscriptionID: to.Ptr("600c95c5-3ee5-44fe-b190-ca38a19adcd7"),
	// 		},
	// }}
}
