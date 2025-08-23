package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/Diagnostics_Update_MaximumSet_Gen.json
func ExampleDiagnosticsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("9D54FE4C-00AF-4836-8F48-B6A9C4E47192", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiagnosticsClient().BeginUpdate(ctx, "rgconfigurationmanager", "testname", armworkloadorchestration.Diagnostic{
		Properties: &armworkloadorchestration.DiagnosticProperties{},
		Tags: map[string]*string{
			"key1922": to.Ptr("efraipifhmdfekwgunngrgvsc"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armworkloadorchestration.DiagnosticsClientUpdateResponse{
	// 	Diagnostic: &armworkloadorchestration.Diagnostic{
	// 		Properties: &armworkloadorchestration.DiagnosticProperties{
	// 			ProvisioningState: to.Ptr(armworkloadorchestration.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armworkloadorchestration.ExtendedLocation{
	// 			Name: to.Ptr("szjrwimeqyiue"),
	// 			Type: to.Ptr(armworkloadorchestration.ExtendedLocationTypeEdgeZone),
	// 		},
	// 		ETag: to.Ptr("oeib"),
	// 		Tags: map[string]*string{
	// 			"key4304": to.Ptr("mdrwpsdrcicagvximokxrrp"),
	// 		},
	// 		Location: to.Ptr("ouwfvnokjvivmjzqpupwrbsmls"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 		Name: to.Ptr("drohk"),
	// 		Type: to.Ptr("kghs"),
	// 		SystemData: &armworkloadorchestration.SystemData{
	// 			CreatedBy: to.Ptr("nvjczgdguyvllp"),
	// 			CreatedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("uzbznzjgvaspvtqhyg"),
	// 			LastModifiedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
	// 		},
	// 	},
	// }
}
