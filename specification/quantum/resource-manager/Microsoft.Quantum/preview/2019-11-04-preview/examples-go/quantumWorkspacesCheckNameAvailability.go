package armquantum_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quantum/armquantum"
)

// x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2019-11-04-preview/examples/quantumWorkspacesCheckNameAvailability.json
func ExampleWorkspaceClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armquantum.NewWorkspaceClient("<subscription-id>", cred, nil)
	res, err := client.CheckNameAvailability(ctx,
		"<location-name>",
		armquantum.CheckNameAvailabilityParameters{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspaceClientCheckNameAvailabilityResult)
}
