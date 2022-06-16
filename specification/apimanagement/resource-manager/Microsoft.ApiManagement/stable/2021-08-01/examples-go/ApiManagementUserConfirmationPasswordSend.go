package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUserConfirmationPasswordSend.json
func ExampleUserConfirmationPasswordClient_Send() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewUserConfirmationPasswordClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Send(ctx,
		"rg1",
		"apimService1",
		"57127d485157a511ace86ae7",
		&armapimanagement.UserConfirmationPasswordClientSendOptions{AppType: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
