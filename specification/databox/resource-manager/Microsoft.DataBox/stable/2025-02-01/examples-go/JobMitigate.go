package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cecf11996f2bfc7958f6d0de814badab23f9720/specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/JobMitigate.json
func ExampleManagementClient_Mitigate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewManagementClient().Mitigate(ctx, "TestJobName1", "YourResourceGroupName", armdatabox.MitigateJobRequest{
		SerialNumberCustomerResolutionMap: map[string]*armdatabox.CustomerResolutionCode{
			"testDISK-1": to.Ptr(armdatabox.CustomerResolutionCodeMoveToCleanUpDevice),
			"testDISK-2": to.Ptr(armdatabox.CustomerResolutionCodeResume),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
