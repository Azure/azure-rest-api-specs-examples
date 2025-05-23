package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/GetSitePrivateLinkResources.json
func ExampleStaticSitesClient_GetPrivateLinkResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStaticSitesClient().GetPrivateLinkResources(ctx, "rg", "testSite", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourcesWrapper = armappservice.PrivateLinkResourcesWrapper{
	// 	Value: []*armappservice.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("site"),
	// 			Type: to.Ptr("Microsoft.Web/sites/privateLinkResources"),
	// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/sites/testSite/privateLinkResources/site"),
	// 			Properties: &armappservice.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("sites"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("sites")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.testsite.azurewebsites.net")},
	// 					},
	// 			}},
	// 		}
}
