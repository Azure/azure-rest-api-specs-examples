package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/685aad3f33d355c1d9c89d493ee9398865367bd8/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Services_CheckNameAvailability.json
func ExampleServicesClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().CheckNameAvailability(ctx, "eastus", armappplatform.NameAvailabilityParameters{
		Name: to.Ptr("myservice"),
		Type: to.Ptr("Microsoft.AppPlatform/Spring"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NameAvailability = armappplatform.NameAvailability{
	// 	Message: to.Ptr("The name is already used."),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr("AlreadyExists"),
	// }
}
