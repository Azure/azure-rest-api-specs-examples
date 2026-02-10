package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
)

// Generated from example definition: 2025-06-01-preview/CheckNameAvailable.json
func ExampleOperationsClient_CheckNameAvailability_configurationStoresCheckNameAvailable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("c80fb759-c965-4c6a-9110-9b2b2d038882", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().CheckNameAvailability(ctx, armappconfiguration.CheckNameAvailabilityParameters{
		Name: to.Ptr("contoso"),
		Type: to.Ptr(armappconfiguration.ConfigurationResourceTypeMicrosoftAppConfigurationConfigurationStores),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappconfiguration.OperationsClientCheckNameAvailabilityResponse{
	// 	NameAvailabilityStatus: &armappconfiguration.NameAvailabilityStatus{
	// 		Message: to.Ptr("The specified name is available."),
	// 		NameAvailable: to.Ptr(true),
	// 	},
	// }
}
