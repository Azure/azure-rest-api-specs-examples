package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/SingleSignOn_CreateOrUpdate_MaximumSet_Gen.json
func ExampleSingleSignOnClient_BeginCreateOrUpdate_singleSignOnCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSingleSignOnClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myMonitor", "default", armdynatrace.SingleSignOnResource{
		Properties: &armdynatrace.SingleSignOnProperties{
			AADDomains: []*string{
				to.Ptr("mpliftrdt20210811outlook.onmicrosoft.com")},
			EnterpriseAppID:   to.Ptr("00000000-0000-0000-0000-000000000000"),
			ProvisioningState: to.Ptr(armdynatrace.ProvisioningStateAccepted),
			SingleSignOnState: to.Ptr(armdynatrace.SingleSignOnStatesEnable),
			SingleSignOnURL:   to.Ptr("https://www.dynatrace.io"),
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
	// res.SingleSignOnResource = armdynatrace.SingleSignOnResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Dynatrace.Observability/monitors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Dynatrace.Observability/monitors/myMonitor/singleSignOnConfigurations/default"),
	// 	Properties: &armdynatrace.SingleSignOnProperties{
	// 		AADDomains: []*string{
	// 			to.Ptr("mpliftrdt20210811outlook.onmicrosoft.com")},
	// 			EnterpriseAppID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			ProvisioningState: to.Ptr(armdynatrace.ProvisioningStateSucceeded),
	// 			SingleSignOnState: to.Ptr(armdynatrace.SingleSignOnStatesEnable),
	// 			SingleSignOnURL: to.Ptr("https://www.dynatrace.io/IAmSomeHash"),
	// 		},
	// 		SystemData: &armdynatrace.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-16T04:25:21.040Z"); return t}()),
	// 			CreatedBy: to.Ptr("alice@microsoft.com"),
	// 			CreatedByType: to.Ptr(armdynatrace.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-16T04:25:21.040Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("alice@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armdynatrace.CreatedByTypeUser),
	// 		},
	// 	}
}
