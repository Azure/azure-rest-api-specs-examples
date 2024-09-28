package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1ad29756bd141a47cac770140105a706d065ae1b/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/ChaosFaultEnableDisable.json
func ExampleChaosFaultClient_BeginEnableDisable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewChaosFaultClient().BeginEnableDisable(ctx, "myResourceGroupName", "myAccountName", "ServiceUnavailability", armcosmos.ChaosFaultResource{
		Properties: &armcosmos.ChaosFaultProperties{
			Action:        to.Ptr(armcosmos.SupportedActionsEnable),
			ContainerName: to.Ptr("testCollection"),
			DatabaseName:  to.Ptr("testDatabase"),
			Region:        to.Ptr("EastUS"),
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
	// res.ChaosFaultResource = armcosmos.ChaosFaultResource{
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/chaosFaults/ServiceUnavailability"),
	// 	Properties: &armcosmos.ChaosFaultProperties{
	// 		Action: to.Ptr(armcosmos.SupportedActionsEnable),
	// 		ContainerName: to.Ptr("testCollection"),
	// 		DatabaseName: to.Ptr("testDatabase"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Region: to.Ptr("EastUS"),
	// 	},
	// }
}
