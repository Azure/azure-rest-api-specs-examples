package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/ListPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_ListByBotResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbotservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().ListByBotResource(ctx, "res6977", "sto2527", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourceListResult = armbotservice.PrivateLinkResourceListResult{
	// 	Value: []*armbotservice.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("resource1"),
	// 			Type: to.Ptr("Microsoft.BotService/botServices/privateLinkResources"),
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.BotService/botServices/sto2527/privateLinkResources/resource1"),
	// 			Properties: &armbotservice.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("bot"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("bot")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.botframework.com")},
	// 					},
	// 			}},
	// 		}
}
