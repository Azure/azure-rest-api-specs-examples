package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/ConfigurationAssignments_GetParent.json
func ExampleConfigurationAssignmentsClient_GetParent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmaintenance.NewConfigurationAssignmentsClient("5b4b650e-28b9-4790-b3ab-ddbd88d727c4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetParent(ctx,
		"examplerg",
		"Microsoft.Compute",
		"virtualMachineScaleSets",
		"smdtest1",
		"virtualMachines",
		"smdvm1",
		"workervmPolicy",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
