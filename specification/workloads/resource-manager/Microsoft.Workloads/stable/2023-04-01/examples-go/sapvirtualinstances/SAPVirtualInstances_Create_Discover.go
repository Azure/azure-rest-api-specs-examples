package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Create_Discover.json
func ExampleSAPVirtualInstancesClient_BeginCreate_registerExistingSapSystemAsVirtualInstanceForSapSolutions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPVirtualInstancesClient().BeginCreate(ctx, "test-rg", "X00", armworkloads.SAPVirtualInstance{
		Location: to.Ptr("northeurope"),
		Tags: map[string]*string{
			"createdby": to.Ptr("abc@microsoft.com"),
			"test":      to.Ptr("abc"),
		},
		Properties: &armworkloads.SAPVirtualInstanceProperties{
			Configuration: &armworkloads.DiscoveryConfiguration{
				ConfigurationType: to.Ptr(armworkloads.SAPConfigurationTypeDiscovery),
				CentralServerVMID: to.Ptr("/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/sapq20scsvm0"),
			},
			Environment: to.Ptr(armworkloads.SAPEnvironmentTypeNonProd),
			SapProduct:  to.Ptr(armworkloads.SAPProductTypeS4HANA),
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
	// res.SAPVirtualInstance = armworkloads.SAPVirtualInstance{
	// 	Name: to.Ptr("Q20"),
	// 	Type: to.Ptr("microsoft.workloads/sapvirtualinstances"),
	// 	ID: to.Ptr("/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/Q20"),
	// 	SystemData: &armworkloads.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-27T08:22:52.3318907Z"); return t}()),
	// 		CreatedBy: to.Ptr("abc@microsoft.com"),
	// 		CreatedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-27T11:44:17.9310503Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("abc@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("northeurope"),
	// 	Properties: &armworkloads.SAPVirtualInstanceProperties{
	// 		Configuration: &armworkloads.DiscoveryConfiguration{
	// 			ConfigurationType: to.Ptr(armworkloads.SAPConfigurationTypeDiscovery),
	// 			AppLocation: to.Ptr("westeurope"),
	// 			CentralServerVMID: to.Ptr("/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg-infra/providers/Microsoft.Compute/virtualMachines/q20ascsvm0"),
	// 		},
	// 		Environment: to.Ptr(armworkloads.SAPEnvironmentTypeNonProd),
	// 		Health: to.Ptr(armworkloads.SAPHealthStateHealthy),
	// 		ManagedResourceGroupConfiguration: &armworkloads.ManagedRGConfiguration{
	// 			Name: to.Ptr("mrg-Q20-5b0097"),
	// 		},
	// 		ProvisioningState: to.Ptr(armworkloads.SapVirtualInstanceProvisioningStateSucceeded),
	// 		SapProduct: to.Ptr(armworkloads.SAPProductTypeS4HANA),
	// 		State: to.Ptr(armworkloads.SAPVirtualInstanceStateRegistrationComplete),
	// 		Status: to.Ptr(armworkloads.SAPVirtualInstanceStatusRunning),
	// 	},
	// }
}
