package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/communication/resource-manager/Microsoft.Communication/preview/2023-06-01-preview/examples/suppressionLists/createOrUpdateSuppressionList.json
func ExampleSuppressionListsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSuppressionListsClient().CreateOrUpdate(ctx, "contosoResourceGroup", "contosoEmailService", "contoso.com", "aaaa1111-bbbb-2222-3333-aaaa11112222", armcommunication.SuppressionListResource{
		Properties: &armcommunication.SuppressionListProperties{
			ListName: to.Ptr("contosoNewsAlerts"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SuppressionListResource = armcommunication.SuppressionListResource{
	// 	Name: to.Ptr("aaaa1111-bbbb-2222-3333-aaaa11112222"),
	// 	Type: to.Ptr("Microsoft.Communication/EmailServices/Domains/SuppressionLists"),
	// 	ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/contosoResourceGroup/providers/Microsoft.Communication/EmailServices/contosoEmailService/Domains/contoso.com/suppressionLists/aaaa1111-bbbb-2222-3333-aaaa11112222"),
	// 	Properties: &armcommunication.SuppressionListProperties{
	// 		CreatedTimeStamp: to.Ptr("2023-06-06T17:06:26.100Z"),
	// 		LastUpdatedTimeStamp: to.Ptr("2023-06-06T17:06:26.100Z"),
	// 		ListName: to.Ptr("contosoNewsAlerts"),
	// 	},
	// }
}
