package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7832c9e47b8998a1c994b98345eea24dbc2ac5b8/specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/suppressionLists/deleteAddress.json
func ExampleSuppressionListAddressesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSuppressionListAddressesClient().Delete(ctx, "MyResourceGroup", "MyEmailServiceResource", "mydomain.com", "aaaa1111-bbbb-2222-3333-aaaa11112222", "11112222-3333-4444-5555-999999999999", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
