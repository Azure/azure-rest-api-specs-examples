package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/ManagedEnvironments_ListWorkloadProfileStates.json
func ExampleManagedEnvironmentsClient_NewListWorkloadProfileStatesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedEnvironmentsClient().NewListWorkloadProfileStatesPager("examplerg", "jlaw-demo1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.WorkloadProfileStatesCollection = armappcontainers.WorkloadProfileStatesCollection{
		// 	Value: []*armappcontainers.WorkloadProfileStates{
		// 		{
		// 			Name: to.Ptr("GP1"),
		// 			Type: to.Ptr("/providers/Microsoft.App/workloadProfileStates"),
		// 			ID: to.Ptr("/subscriptions/55f240e3-3d66-44f6-8358-4e4f3d7a2e51/providers/Microsoft.App/workloadProfileStates/GP1"),
		// 			Properties: &armappcontainers.WorkloadProfileStatesProperties{
		// 				CurrentCount: to.Ptr[int32](3),
		// 				MaximumCount: to.Ptr[int32](10),
		// 				MinimumCount: to.Ptr[int32](3),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MO3"),
		// 			Type: to.Ptr("/providers/Microsoft.App/workloadProfileStates"),
		// 			ID: to.Ptr("/subscriptions/55f240e3-3d66-44f6-8358-4e4f3d7a2e51/providers/Microsoft.App/workloadProfileStates/MO3"),
		// 			Properties: &armappcontainers.WorkloadProfileStatesProperties{
		// 				CurrentCount: to.Ptr[int32](0),
		// 				MaximumCount: to.Ptr[int32](2),
		// 				MinimumCount: to.Ptr[int32](0),
		// 			},
		// 	}},
		// }
	}
}
