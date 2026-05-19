package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/CdnProfiles_Migrate.json
func ExampleProfilesClient_BeginCdnMigrateToAfd() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProfilesClient().BeginCdnMigrateToAfd(ctx, "RG", "profile1", armcdn.MigrationToAfdParameters{
		SKU: &armcdn.SKU{
			Name: to.Ptr(armcdn.SKUNameStandardAzureFrontDoor),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcdn.ProfilesClientCdnMigrateToAfdResponse{
	// 	MigrateResult: armcdn.MigrateResult{
	// 		Type: to.Ptr("Microsoft.Cdn/profiles/migrate"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/RG/providers/Microsoft.Cdn/operationresults/operationid/profileresults/profile1/migrateresults/profile1"),
	// 		Properties: &armcdn.MigrateResultProperties{
	// 			MigratedProfileResourceID: &armcdn.ResourceReference{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1"),
	// 			},
	// 		},
	// 	},
	// }
}
