package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClusterPrincipalAssignmentsCheckNameAvailability.json
func ExampleClusterPrincipalAssignmentsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewClusterPrincipalAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.CheckNameAvailability(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armkusto.ClusterPrincipalAssignmentCheckNameRequest{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClusterPrincipalAssignmentsClientCheckNameAvailabilityResult)
}
