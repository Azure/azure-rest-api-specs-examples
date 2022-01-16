Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmmanagedapplications%2Fv0.2.0/sdk/resourcemanager/resources/armmanagedapplications/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagedapplications_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/createOrUpdateApplicationDefinition.json
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
			Location: to.StringPtr("<location>"),
			Properties: &armmanagedapplications.ApplicationDefinitionProperties{
				Description: to.StringPtr("<description>"),
				Authorizations: []*armmanagedapplications.ApplicationProviderAuthorization{
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
