package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v3"
)

// Generated from example definition: 2025-09-01/communicationServices/checkNameAvailabilityAvailable.json
func ExampleServicesClient_CheckNameAvailability_checkNameAvailabilityAvailable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("11112222-3333-4444-5555-666677778888", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().CheckNameAvailability(ctx, armcommunication.NameAvailabilityParameters{
		Name: to.Ptr("MyCommunicationService"),
		Type: to.Ptr("Microsoft.Communication/CommunicationServices"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcommunication.ServicesClientCheckNameAvailabilityResponse{
	// 	CheckNameAvailabilityResponse: &armcommunication.CheckNameAvailabilityResponse{
	// 		Message: to.Ptr("Requested name is available for the requested type"),
	// 		NameAvailable: to.Ptr(true),
	// 		Reason: to.Ptr(armcommunication.CheckNameAvailabilityReason("NameAvailable")),
	// 	},
	// }
}
