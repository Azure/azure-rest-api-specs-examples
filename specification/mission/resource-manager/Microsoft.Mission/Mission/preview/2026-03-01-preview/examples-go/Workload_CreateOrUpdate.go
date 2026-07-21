package armenclave_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/enclave/armenclave"
)

// Generated from example definition: 2026-03-01-preview/Workload_CreateOrUpdate.json
func ExampleWorkloadClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armenclave.NewClientFactory("CA1CB369-DD26-4DB2-9D43-9AFEF0F22093", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadClient().BeginCreateOrUpdate(ctx, "rgopenapi", "TestMyEnclave", "TestMyWorkload", armenclave.WorkloadResource{
		Properties: &armenclave.WorkloadProperties{
			ResourceGroupCollection: []*string{},
		},
		Tags: map[string]*string{
			"TestKey": to.Ptr("TestValue"),
		},
		Location: to.Ptr("westcentralus"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armenclave.WorkloadClientCreateOrUpdateResponse{
	// 	WorkloadResource: armenclave.WorkloadResource{
	// 		Properties: &armenclave.WorkloadProperties{
	// 			ProvisioningState: to.Ptr(armenclave.ProvisioningStateSucceeded),
	// 			ResourceGroupCollection: []*string{
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"TestKey": to.Ptr("TestValue"),
	// 		},
	// 		Location: to.Ptr("westcentralus"),
	// 		ID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/microsoft.mission/virtualenclaves/TestMyEnclave/workloads/kxzylwqnmxtivpmupnlho"),
	// 		Name: to.Ptr("kxzylwqnmxtivpmupnlho"),
	// 		Type: to.Ptr("microsoft.mission/virtualenclaves/workloads"),
	// 		SystemData: &armenclave.SystemData{
	// 			CreatedBy: to.Ptr("myAlias"),
	// 			CreatedByType: to.Ptr(armenclave.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("myAlias"),
	// 			LastModifiedByType: to.Ptr(armenclave.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 		},
	// 	},
	// }
}
