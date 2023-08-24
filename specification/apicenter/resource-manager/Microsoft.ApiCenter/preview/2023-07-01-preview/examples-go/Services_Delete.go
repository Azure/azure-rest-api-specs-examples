package armapicenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apicenter/armapicenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apicenter/resource-manager/Microsoft.ApiCenter/preview/2023-07-01-preview/examples/Services_Delete.json
func ExampleServicesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("<subscription-id>", "<service-name>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewServicesClient().Delete(ctx, "contoso-resources", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
