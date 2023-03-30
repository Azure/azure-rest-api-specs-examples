package armcustomproviders_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customproviders/armcustomproviders"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/customproviders/resource-manager/Microsoft.CustomProviders/preview/2018-09-01-preview/examples/createOrUpdateAssociation.json
func ExampleAssociationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomproviders.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssociationsClient().BeginCreateOrUpdate(ctx, "scope", "associationName", armcustomproviders.Association{
		Properties: &armcustomproviders.AssociationProperties{
			TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/appRG/providers/Microsoft.Solutions/applications/applicationName"),
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
	// res.Association = armcustomproviders.Association{
	// 	Name: to.Ptr("associationName"),
	// 	Type: to.Ptr("Microsoft.CustomProviders/associations"),
	// 	ID: to.Ptr("/scope/providers/Microsoft.CustomProviders/associations/associationName"),
	// 	Properties: &armcustomproviders.AssociationProperties{
	// 		ProvisioningState: to.Ptr(armcustomproviders.ProvisioningStateSucceeded),
	// 		TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/appRG/providers/Microsoft.Solutions/applications/applicationName"),
	// 	},
	// }
}
