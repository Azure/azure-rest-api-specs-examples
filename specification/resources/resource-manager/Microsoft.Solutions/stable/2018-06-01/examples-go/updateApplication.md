Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmmanagedapplications%2Fv0.4.0/sdk/resourcemanager/resources/armmanagedapplications/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/updateApplication.json
func ExampleApplicationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmanagedapplications.NewApplicationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<application-name>",
		&armmanagedapplications.ApplicationsClientUpdateOptions{Parameters: &armmanagedapplications.ApplicationPatchable{
			Kind: to.Ptr("<kind>"),
			Properties: &armmanagedapplications.ApplicationPropertiesPatchable{
				ApplicationDefinitionID: to.Ptr("<application-definition-id>"),
				ManagedResourceGroupID:  to.Ptr("<managed-resource-group-id>"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
