Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmmanagedapplications%2Fv1.0.0/sdk/resourcemanager/resources/armmanagedapplications/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/createOrUpdateApplicationDefinition.json
func ExampleApplicationDefinitionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagedapplications.NewApplicationDefinitionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg",
		"myManagedApplicationDef",
		armmanagedapplications.ApplicationDefinition{
			Location: to.Ptr("East US 2"),
			Properties: &armmanagedapplications.ApplicationDefinitionProperties{
				Description: to.Ptr("myManagedApplicationDef description"),
				Authorizations: []*armmanagedapplications.ApplicationProviderAuthorization{
					{
						PrincipalID:      to.Ptr("validprincipalguid"),
						RoleDefinitionID: to.Ptr("validroleguid"),
					}},
				DisplayName:    to.Ptr("myManagedApplicationDef"),
				LockLevel:      to.Ptr(armmanagedapplications.ApplicationLockLevelNone),
				PackageFileURI: to.Ptr("https://path/to/packagezipfile"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
