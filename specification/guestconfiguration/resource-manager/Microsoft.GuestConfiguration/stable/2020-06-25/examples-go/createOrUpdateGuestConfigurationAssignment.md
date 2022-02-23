Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fguestconfiguration%2Farmguestconfiguration%2Fv0.2.1/sdk/resourcemanager/guestconfiguration/armguestconfiguration/README.md) on how to add the SDK to your project and authenticate.

```go
package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
)

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/createOrUpdateGuestConfigurationAssignment.json
func ExampleAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<guest-configuration-assignment-name>",
		"<resource-group-name>",
		"<vm-name>",
		armguestconfiguration.Assignment{
			Name:     to.StringPtr("<name>"),
			Location: to.StringPtr("<location>"),
			Properties: &armguestconfiguration.AssignmentProperties{
				Context: to.StringPtr("<context>"),
				GuestConfiguration: &armguestconfiguration.Navigation{
					Name:           to.StringPtr("<name>"),
					AssignmentType: armguestconfiguration.AssignmentType("ApplyAndAutoCorrect").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.AssignmentsClientCreateOrUpdateResult)
}
```
