package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/communicationServices/checkNameAvailabilityAvailable.json
func ExampleServicesClient_CheckNameAvailability_checkNameAvailabilityAvailable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcommunication.NewServicesClient("12345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckNameAvailability(ctx, armcommunication.NameAvailabilityParameters{
		Name: to.Ptr("MyCommunicationService"),
		Type: to.Ptr("Microsoft.Communication/CommunicationServices"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
