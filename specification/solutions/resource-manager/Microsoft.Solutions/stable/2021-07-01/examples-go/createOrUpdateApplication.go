package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateApplication.json
func ExampleApplicationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginCreateOrUpdate(ctx, "rg", "myManagedApplication", armmanagedapplications.Application{
		Kind: to.Ptr("ServiceCatalog"),
		Properties: &armmanagedapplications.ApplicationProperties{
			ApplicationDefinitionID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef"),
			ManagedResourceGroupID:  to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
		},
	}, nil)
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
	// res.Application = armmanagedapplications.Application{
	// 	Name: to.Ptr("myManagedApplication"),
	// 	Type: to.Ptr("Microsoft.Solutions/applications"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication"),
	// 	Kind: to.Ptr("ServiceCatalog"),
	// 	Properties: &armmanagedapplications.ApplicationProperties{
	// 		ApplicationDefinitionID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef"),
	// 		ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
	// 		ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
	// 	},
	// }
}
