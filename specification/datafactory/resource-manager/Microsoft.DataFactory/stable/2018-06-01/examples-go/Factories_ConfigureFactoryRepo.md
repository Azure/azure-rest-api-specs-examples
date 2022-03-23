Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatafactory%2Farmdatafactory%2Fv0.3.0/sdk/resourcemanager/datafactory/armdatafactory/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_ConfigureFactoryRepo.json
func ExampleFactoriesClient_ConfigureFactoryRepo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewFactoriesClient("<subscription-id>", cred, nil)
	res, err := client.ConfigureFactoryRepo(ctx,
		"<location-id>",
		armdatafactory.FactoryRepoUpdate{
			FactoryResourceID: to.StringPtr("<factory-resource-id>"),
			RepoConfiguration: &armdatafactory.FactoryVSTSConfiguration{
				Type:                to.StringPtr("<type>"),
				AccountName:         to.StringPtr("<account-name>"),
				CollaborationBranch: to.StringPtr("<collaboration-branch>"),
				LastCommitID:        to.StringPtr("<last-commit-id>"),
				RepositoryName:      to.StringPtr("<repository-name>"),
				RootFolder:          to.StringPtr("<root-folder>"),
				ProjectName:         to.StringPtr("<project-name>"),
				TenantID:            to.StringPtr("<tenant-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FactoriesClientConfigureFactoryRepoResult)
}
```
