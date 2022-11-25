package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_ConfigureFactoryRepo.json
func ExampleFactoriesClient_ConfigureFactoryRepo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewFactoriesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ConfigureFactoryRepo(ctx, "East US", armdatafactory.FactoryRepoUpdate{
		FactoryResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
		RepoConfiguration: &armdatafactory.FactoryVSTSConfiguration{
			Type:                to.Ptr("FactoryVSTSConfiguration"),
			AccountName:         to.Ptr("ADF"),
			CollaborationBranch: to.Ptr("master"),
			LastCommitID:        to.Ptr(""),
			RepositoryName:      to.Ptr("repo"),
			RootFolder:          to.Ptr("/"),
			ProjectName:         to.Ptr("project"),
			TenantID:            to.Ptr(""),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
