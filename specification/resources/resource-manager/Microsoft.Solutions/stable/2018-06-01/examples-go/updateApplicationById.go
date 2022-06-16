package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/updateApplicationById.json
func ExampleApplicationsClient_UpdateByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedapplications.NewApplicationsClient("<subscription-id>", cred, nil)
	res, err := client.UpdateByID(ctx,
		"<application-id>",
		&armmanagedapplications.ApplicationsUpdateByIDOptions{Parameters: &armmanagedapplications.Application{
			Kind: to.StringPtr("<kind>"),
			Properties: &armmanagedapplications.ApplicationProperties{
				ApplicationDefinitionID: to.StringPtr("<application-definition-id>"),
				ManagedResourceGroupID:  to.StringPtr("<managed-resource-group-id>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Application.ID: %s\n", *res.ID)
}
