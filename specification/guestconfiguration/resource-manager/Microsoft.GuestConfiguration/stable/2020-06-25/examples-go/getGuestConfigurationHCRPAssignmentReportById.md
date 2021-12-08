Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fguestconfiguration%2Farmguestconfiguration%2Fv0.1.0/sdk/resourcemanager/guestconfiguration/armguestconfiguration/README.md) on how to add the SDK to your project and authenticate.

```go
package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
)

// x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2020-06-25/examples/getGuestConfigurationHCRPAssignmentReportById.json
func ExampleGuestConfigurationHCRPAssignmentReportsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armguestconfiguration.NewGuestConfigurationHCRPAssignmentReportsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<guest-configuration-assignment-name>",
		"<report-id>",
		"<machine-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GuestConfigurationAssignmentReport.ID: %s\n", *res.ID)
}
```
