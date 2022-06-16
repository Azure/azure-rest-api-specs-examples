package armaddons_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/addons/armaddons"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/addons/resource-manager/Microsoft.Addons/preview/2018-03-01/examples/SupportPlanTypes_Delete.json
func ExampleSupportPlanTypesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armaddons.NewSupportPlanTypesClient("d18d258f-bdba-4de1-8b51-e79d6c181d5e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"Canonical",
		armaddons.PlanTypeNameStandard,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
