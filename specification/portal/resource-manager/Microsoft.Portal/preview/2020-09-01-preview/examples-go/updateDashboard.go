package armportal_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/updateDashboard.json
func ExampleDashboardsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armportal.NewDashboardsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"testRG",
		"testDashboard",
		armportal.PatchableDashboard{
			Tags: map[string]*string{
				"aKey":       to.Ptr("bValue"),
				"anotherKey": to.Ptr("anotherValue2"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
