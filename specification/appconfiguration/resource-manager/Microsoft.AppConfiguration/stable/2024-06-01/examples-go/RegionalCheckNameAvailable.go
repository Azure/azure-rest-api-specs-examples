package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-06-01/examples/RegionalCheckNameAvailable.json
func ExampleOperationsClient_RegionalCheckNameAvailability_configurationStoresCheckNameAvailable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().RegionalCheckNameAvailability(ctx, "westus", armappconfiguration.CheckNameAvailabilityParameters{
		Name: to.Ptr("contoso"),
		Type: to.Ptr(armappconfiguration.ConfigurationResourceTypeMicrosoftAppConfigurationConfigurationStores),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NameAvailabilityStatus = armappconfiguration.NameAvailabilityStatus{
	// 	Message: to.Ptr("The specified name is available."),
	// 	NameAvailable: to.Ptr(true),
	// }
}
