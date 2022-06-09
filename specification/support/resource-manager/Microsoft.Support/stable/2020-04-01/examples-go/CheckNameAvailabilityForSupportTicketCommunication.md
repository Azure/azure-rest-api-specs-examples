```go
package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/CheckNameAvailabilityForSupportTicketCommunication.json
func ExampleCommunicationsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsupport.NewCommunicationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckNameAvailability(ctx,
		"testticket",
		armsupport.CheckNameAvailabilityInput{
			Name: to.Ptr("sampleName"),
			Type: to.Ptr(armsupport.TypeMicrosoftSupportCommunications),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsupport%2Farmsupport%2Fv1.0.0/sdk/resourcemanager/support/armsupport/README.md) on how to add the SDK to your project and authenticate.
