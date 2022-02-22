Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsolutions%2Farmmanagedapplications%2Fv0.2.1/sdk/resourcemanager/solutions/armmanagedapplications/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagedapplications_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications"
)

// x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateApplicationDefinition.json
func ExampleApplicationDefinitionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedapplications.NewApplicationDefinitionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<application-definition-name>",
		armmanagedapplications.ApplicationDefinition{
			Properties: &armmanagedapplications.ApplicationDefinitionProperties{
				Description: to.StringPtr("<description>"),
				Authorizations: []*armmanagedapplications.ApplicationAuthorization{
					{
						PrincipalID:      to.StringPtr("<principal-id>"),
						RoleDefinitionID: to.StringPtr("<role-definition-id>"),
					}},
				DisplayName:    to.StringPtr("<display-name>"),
				LockLevel:      armmanagedapplications.ApplicationLockLevelNone.ToPtr(),
				PackageFileURI: to.StringPtr("<package-file-uri>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ApplicationDefinitionsClientCreateOrUpdateResult)
}
```
