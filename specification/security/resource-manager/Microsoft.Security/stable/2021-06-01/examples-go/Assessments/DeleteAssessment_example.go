package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/DeleteAssessment_example.json
func ExampleAssessmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAssessmentsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2",
		"8bb8be0a-6010-4789-812f-e4d661c4ed0e",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
