Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foperationsmanagement%2Farmoperationsmanagement%2Fv0.3.1/sdk/resourcemanager/operationsmanagement/armoperationsmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armoperationsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationsmanagement/armoperationsmanagement"
)

// x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/ManagementAssociationCreate.json
func ExampleManagementAssociationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armoperationsmanagement.NewManagementAssociationsClient("<subscription-id>",
		"<provider-name>",
		"<resource-type>",
		"<resource-name>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<management-association-name>",
		armoperationsmanagement.ManagementAssociation{
			Location: to.StringPtr("<location>"),
			Properties: &armoperationsmanagement.ManagementAssociationProperties{
				ApplicationID: to.StringPtr("<application-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ManagementAssociationsClientCreateOrUpdateResult)
}
```
