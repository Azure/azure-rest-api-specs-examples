package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_Migrate.json
func ExampleProfilesClient_BeginMigrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProfilesClient().BeginMigrate(ctx, "RG", armcdn.MigrationParameters{
		ClassicResourceReference: &armcdn.ResourceReference{
			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoors/frontdoorname"),
		},
		ProfileName: to.Ptr("profile1"),
		SKU: &armcdn.SKU{
			Name: to.Ptr(armcdn.SKUNameStandardAzureFrontDoor),
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
	// res.MigrateResult = armcdn.MigrateResult{
	// 	Type: to.Ptr("Microsoft.Cdn/migrate"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/operationresults/operationid/profileresults/profile1/migrateresults/profile1"),
	// 	Properties: &armcdn.MigrateResultProperties{
	// 		MigratedProfileResourceID: &armcdn.ResourceReference{
	// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1"),
	// 		},
	// 	},
	// }
}
