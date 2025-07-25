package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/ListDeploymentSkus.json
func ExampleDeploymentsClient_NewListSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeploymentsClient().NewListSKUsPager("resourceGroupName", "accountName", "deploymentName", nil)
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
		// page.DeploymentSKUListResult = armcognitiveservices.DeploymentSKUListResult{
		// 	Value: []*armcognitiveservices.SKUResource{
		// 		{
		// 			Capacity: &armcognitiveservices.CapacityConfig{
		// 				Default: to.Ptr[int32](100),
		// 				AllowedValues: []*int32{
		// 					to.Ptr[int32](100),
		// 					to.Ptr[int32](200)},
		// 					Maximum: to.Ptr[int32](1000),
		// 					Minimum: to.Ptr[int32](100),
		// 					Step: to.Ptr[int32](100),
		// 				},
		// 				ResourceType: to.Ptr("Microsoft.CognitiveServices/accounts/deployments"),
		// 				SKU: &armcognitiveservices.SKU{
		// 					Name: to.Ptr("Standard"),
		// 					Capacity: to.Ptr[int32](1),
		// 				},
		// 		}},
		// 	}
	}
}
