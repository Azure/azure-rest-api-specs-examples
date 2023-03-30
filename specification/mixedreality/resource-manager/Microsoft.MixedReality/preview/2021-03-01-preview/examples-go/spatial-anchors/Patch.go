package armmixedreality_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/spatial-anchors/Patch.json
func ExampleSpatialAnchorsAccountsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmixedreality.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSpatialAnchorsAccountsClient().Update(ctx, "MyResourceGroup", "MyAccount", armmixedreality.SpatialAnchorsAccount{
		Location: to.Ptr("eastus2euap"),
		Tags: map[string]*string{
			"hero":    to.Ptr("romeo"),
			"heroine": to.Ptr("juliet"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SpatialAnchorsAccount = armmixedreality.SpatialAnchorsAccount{
	// 	Name: to.Ptr("MyAccount"),
	// 	Type: to.Ptr("Microsoft.MixedReality/spatialAnchorsAccounts"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/MyResourceGroup/providers/Microsoft.MixedReality/spatialAnchorsAccounts/MyAccount"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Tags: map[string]*string{
	// 		"hero": to.Ptr("romeo"),
	// 		"heroine": to.Ptr("juliet"),
	// 	},
	// 	Properties: &armmixedreality.AccountProperties{
	// 		AccountDomain: to.Ptr("mixedreality.azure.com"),
	// 		AccountID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	// 	},
	// }
}
