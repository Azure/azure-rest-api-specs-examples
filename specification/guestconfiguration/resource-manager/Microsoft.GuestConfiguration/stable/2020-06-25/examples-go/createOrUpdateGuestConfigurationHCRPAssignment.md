Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fguestconfiguration%2Farmguestconfiguration%2Fv0.4.0/sdk/resourcemanager/guestconfiguration/armguestconfiguration/README.md) on how to add the SDK to your project and authenticate.

```go
package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/createOrUpdateGuestConfigurationHCRPAssignment.json
func ExampleHCRPAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armguestconfiguration.NewHCRPAssignmentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<guest-configuration-assignment-name>",
		"<resource-group-name>",
		"<machine-name>",
		armguestconfiguration.Assignment{
			Name:     to.Ptr("<name>"),
			Location: to.Ptr("<location>"),
			Properties: &armguestconfiguration.AssignmentProperties{
				Context: to.Ptr("<context>"),
				GuestConfiguration: &armguestconfiguration.Navigation{
					Name:           to.Ptr("<name>"),
					AssignmentType: to.Ptr(armguestconfiguration.AssignmentTypeApplyAndAutoCorrect),
					ConfigurationParameter: []*armguestconfiguration.ConfigurationParameter{
						{
							Name:  to.Ptr("<name>"),
							Value: to.Ptr("<value>"),
						}},
					ContentHash: to.Ptr("<content-hash>"),
					ContentURI:  to.Ptr("<content-uri>"),
					Version:     to.Ptr("<version>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
