package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/CheckinManifest.json
func ExampleClient_CheckinManifest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().CheckinManifest(ctx, "Microsoft.Contoso", armproviderhub.CheckinManifestParams{
		BaselineArmManifestLocation: to.Ptr("EastUS2EUAP"),
		Environment:                 to.Ptr("Prod"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armproviderhub.ClientCheckinManifestResponse{
	// 	CheckinManifestInfo: armproviderhub.CheckinManifestInfo{
	// 		IsCheckedIn: to.Ptr(false),
	// 		StatusMessage: to.Ptr("Manifest is successfully merged. Use the Default/Custom rollout (http://aka.ms/rpaasrollout) to roll out the manifest in ARM."),
	// 	},
	// }
}
