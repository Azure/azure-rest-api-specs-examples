package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/CheckNameAvailability_NotAvailable.json
func ExampleClient_CheckNameAvailability_availabilityStatusOfResourceNameWhenResourceWithSameNameTypeAndSubscriptionExists() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().CheckNameAvailability(ctx, "resGroupBar", "westus", armrecoveryservices.CheckNameAvailabilityParameters{
		Name: to.Ptr("swaggerExample2"),
		Type: to.Ptr("Microsoft.RecoveryServices/Vaults"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResult = armrecoveryservices.CheckNameAvailabilityResult{
	// 	Message: to.Ptr("Resource already exists with the same name."),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr("AlreadyExists"),
	// }
}
