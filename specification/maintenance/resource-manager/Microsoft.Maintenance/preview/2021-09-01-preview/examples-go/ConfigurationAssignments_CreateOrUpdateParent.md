Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmaintenance%2Farmmaintenance%2Fv0.1.0/sdk/resourcemanager/maintenance/armmaintenance/README.md) on how to add the SDK to your project and authenticate.

```go
package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/ConfigurationAssignments_CreateOrUpdateParent.json
func ExampleConfigurationAssignmentsClient_CreateOrUpdateParent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaintenance.NewConfigurationAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdateParent(ctx,
		"<resource-group-name>",
		"<provider-name>",
		"<resource-parent-type>",
		"<resource-parent-name>",
		"<resource-type>",
		"<resource-name>",
		"<configuration-assignment-name>",
		armmaintenance.ConfigurationAssignment{
			Properties: &armmaintenance.ConfigurationAssignmentProperties{
				MaintenanceConfigurationID: to.StringPtr("<maintenance-configuration-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationAssignment.ID: %s\n", *res.ID)
}
```
