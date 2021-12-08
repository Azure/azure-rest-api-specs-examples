Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fguestconfiguration%2Farmguestconfiguration%2Fv0.1.0/sdk/resourcemanager/guestconfiguration/armguestconfiguration/README.md) on how to add the SDK to your project and authenticate.

```go
package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
)

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/createOrUpdateGuestConfigurationHCRPAssignment.json
func ExampleGuestConfigurationHCRPAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewGuestConfigurationHCRPAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<guest-configuration-assignment-name>",
		"<resource-group-name>",
		"<machine-name>",
		armguestconfiguration.GuestConfigurationAssignment{
			ProxyResource: armguestconfiguration.ProxyResource{
				Resource: armguestconfiguration.Resource{
					Name:     to.StringPtr("<name>"),
					Location: to.StringPtr("<location>"),
				},
			},
			Properties: &armguestconfiguration.GuestConfigurationAssignmentProperties{
				Context: to.StringPtr("<context>"),
				GuestConfiguration: &armguestconfiguration.GuestConfigurationNavigation{
					Name:           to.StringPtr("<name>"),
					AssignmentType: armguestconfiguration.AssignmentTypeApplyAndAutoCorrect.ToPtr(),
					ConfigurationParameter: []*armguestconfiguration.ConfigurationParameter{
						{
							Name:  to.StringPtr("<name>"),
							Value: to.StringPtr("<value>"),
						}},
					ContentHash: to.StringPtr("<content-hash>"),
					ContentURI:  to.StringPtr("<content-uri>"),
					Version:     to.StringPtr("<version>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GuestConfigurationAssignment.ID: %s\n", *res.ID)
}
```
