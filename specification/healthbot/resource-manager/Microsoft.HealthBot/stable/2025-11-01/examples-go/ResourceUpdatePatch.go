package armhealthbot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot/v2"
)

// Generated from example definition: 2025-11-01/ResourceUpdatePatch.json
func ExampleBotsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthbot.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBotsClient().BeginUpdate(ctx, "healthbotClient", "samplebotname", armhealthbot.UpdateParameters{
		SKU: &armhealthbot.SKU{
			Name: to.Ptr(armhealthbot.SKUNameF0),
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
	// res = armhealthbot.BotsClientUpdateResponse{
	// 	HealthBot: &armhealthbot.HealthBot{
	// 		Name: to.Ptr("samplebotname"),
	// 		Type: to.Ptr("Microsoft.HealthBot/healthBots"),
	// 		ID: to.Ptr("/subscriptions/subscription-id/resourceGroups/OneResourceGroupName/providers/Microsoft.HealthBot/healthBots/samplebotname"),
	// 		Identity: &armhealthbot.Identity{
	// 			Type: to.Ptr(armhealthbot.ResourceIdentityTypeSystemAssignedUserAssigned),
	// 			PrincipalID: to.Ptr("principalId"),
	// 			TenantID: to.Ptr("tenantId"),
	// 			UserAssignedIdentities: map[string]*armhealthbot.UserAssignedIdentity{
	// 				"/subscriptions/subscription-id/resourcegroups/myrg/providers/microsoft.managedidentity/userassignedidentities/my-mi": &armhealthbot.UserAssignedIdentity{
	// 				},
	// 				"/subscriptions/subscription-id/resourcegroups/myrg/providers/microsoft.managedidentity/userassignedidentities/my-mi2": &armhealthbot.UserAssignedIdentity{
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armhealthbot.Properties{
	// 			BotManagementPortalLink: to.Ptr("https://us.healthbot.microsoft.com/account/contoso"),
	// 		},
	// 		SKU: &armhealthbot.SKU{
	// 			Name: to.Ptr(armhealthbot.SKUNameF0),
	// 		},
	// 		SystemData: &armhealthbot.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-05T17:18:19.1234567Z"); return t}()),
	// 			CreatedBy: to.Ptr("jack@outlook.com"),
	// 			CreatedByType: to.Ptr(armhealthbot.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-06T17:18:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ryan@outlook.com"),
	// 			LastModifiedByType: to.Ptr(armhealthbot.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
