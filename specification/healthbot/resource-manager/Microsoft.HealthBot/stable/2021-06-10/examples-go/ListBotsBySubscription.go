package armhealthbot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2021-06-10/examples/ListBotsBySubscription.json
func ExampleBotsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthbot.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBotsClient().NewListPager(nil)
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
		// page.BotResponseList = armhealthbot.BotResponseList{
		// 	Value: []*armhealthbot.HealthBot{
		// 		{
		// 			Name: to.Ptr("samplebotname2"),
		// 			Type: to.Ptr("Microsoft.HealthBot/healthBots"),
		// 			ID: to.Ptr("/subscriptions/subscription-id/resourceGroups/OneResourceGroupName/providers/Microsoft.HealthBot/healthBots/samplebotname2"),
		// 			SystemData: &armhealthbot.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-05T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("jack@outlook.com"),
		// 				CreatedByType: to.Ptr(armhealthbot.IdentityTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-06T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("ryan@outlook.com"),
		// 				LastModifiedByType: to.Ptr(armhealthbot.IdentityTypeUser),
		// 			},
		// 			Location: to.Ptr("East US"),
		// 			Identity: &armhealthbot.Identity{
		// 				Type: to.Ptr(armhealthbot.ResourceIdentityTypeSystemAssignedUserAssigned),
		// 				PrincipalID: to.Ptr("principalId"),
		// 				TenantID: to.Ptr("tenantId"),
		// 				UserAssignedIdentities: map[string]*armhealthbot.UserAssignedIdentity{
		// 					"/subscriptions/subscription-id/resourcegroups/myrg/providers/microsoft.managedidentity/userassignedidentities/my-mi": &armhealthbot.UserAssignedIdentity{
		// 					},
		// 					"/subscriptions/subscription-id/resourcegroups/myrg/providers/microsoft.managedidentity/userassignedidentities/my-mi2": &armhealthbot.UserAssignedIdentity{
		// 					},
		// 				},
		// 			},
		// 			Properties: &armhealthbot.Properties{
		// 				BotManagementPortalLink: to.Ptr("https://us.healthbot.microsoft.com/account/samplebotname2-hdi1osc"),
		// 			},
		// 			SKU: &armhealthbot.SKU{
		// 				Name: to.Ptr(armhealthbot.SKUNameS1),
		// 			},
		// 	}},
		// }
	}
}
