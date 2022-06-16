package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateCloud.json
func ExampleCloudsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armscvmm.NewCloudsClient("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testrg",
		"HRCloud",
		armscvmm.Cloud{
			ExtendedLocation: &armscvmm.ExtendedLocation{
				Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
				Type: to.Ptr("customLocation"),
			},
			Location: to.Ptr("East US"),
			Properties: &armscvmm.CloudProperties{
				UUID:        to.Ptr("aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
				VmmServerID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
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
