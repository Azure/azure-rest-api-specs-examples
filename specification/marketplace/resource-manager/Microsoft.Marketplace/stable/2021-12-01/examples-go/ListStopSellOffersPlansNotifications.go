package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/ListStopSellOffersPlansNotifications.json
func ExamplePrivateStoreClient_ListStopSellOffersPlansNotifications() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreClient().ListStopSellOffersPlansNotifications(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", &armmarketplace.PrivateStoreClientListStopSellOffersPlansNotificationsOptions{StopSellSubscriptions: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StopSellOffersPlansNotificationsList = armmarketplace.StopSellOffersPlansNotificationsList{
	// 	StopSellNotifications: []*armmarketplace.StopSellOffersPlansNotificationsListProperties{
	// 		{
	// 			DisplayName: to.Ptr("Offer display name"),
	// 			Icon: to.Ptr("https://some-images.someDomail.com/image/apps.12345678-4321"),
	// 			IsEntire: to.Ptr(false),
	// 			MessageCode: to.Ptr[int64](10002),
	// 			OfferID: to.Ptr("publisherId.legacyId"),
	// 			Plans: []*armmarketplace.PlanNotificationDetails{
	// 				{
	// 					PlanDisplayName: to.Ptr("test123"),
	// 					PlanID: to.Ptr("aaa"),
	// 			}},
	// 			PublicContext: to.Ptr(false),
	// 			SubscriptionsIDs: []*string{
	// 				to.Ptr("090dcec8-3415-4e13-9377-07f489cdfeed")},
	// 		}},
	// 	}
}
