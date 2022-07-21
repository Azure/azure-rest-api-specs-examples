package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/stable/2022-02-22/examples/updateHybridRunbookWorkerGroup.json
func ExampleHybridRunbookWorkerGroupClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomation.NewHybridRunbookWorkerGroupClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"rg",
		"testaccount",
		"TestHybridGroup",
		armautomation.HybridRunbookWorkerGroupCreateOrUpdateParameters{
			Properties: &armautomation.HybridRunbookWorkerGroupCreateOrUpdateProperties{
				Credential: &armautomation.RunAsCredentialAssociationProperty{
					Name: to.Ptr("myRunAsCredentialUpdatedName"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
