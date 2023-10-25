package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Profiles_Update.json
func ExampleProfilesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProfilesClient().BeginUpdate(ctx, "RG", "profile1", armcdn.ProfileUpdateParameters{
		Tags: map[string]*string{
			"additionalProperties": to.Ptr("Tag1"),
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
	// res.Profile = armcdn.Profile{
	// 	Name: to.Ptr("profile1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1"),
	// 	Location: to.Ptr("global"),
	// 	Tags: map[string]*string{
	// 		"additionalProperties": to.Ptr("Tag1"),
	// 	},
	// 	Kind: to.Ptr("frontdoor"),
	// 	Properties: &armcdn.ProfileProperties{
	// 		FrontDoorID: to.Ptr("3b4682da-b3e2-47a1-96ca-08ab3cb7294e"),
	// 		OriginResponseTimeoutSeconds: to.Ptr[int32](30),
	// 		ProvisioningState: to.Ptr(armcdn.ProfileProvisioningStateSucceeded),
	// 		ResourceState: to.Ptr(armcdn.ProfileResourceStateActive),
	// 	},
	// 	SKU: &armcdn.SKU{
	// 		Name: to.Ptr(armcdn.SKUNamePremiumAzureFrontDoor),
	// 	},
	// }
}
