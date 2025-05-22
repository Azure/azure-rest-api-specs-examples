package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: 2024-09-01/SapVirtualInstances_CreateRegisterSapSolutions.json
func ExampleSAPVirtualInstancesClient_BeginCreate_registerExistingSapSystemAsVirtualInstanceForSapSolutions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("8e17e36c-42e9-4cd5-a078-7b44883414e0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPVirtualInstancesClient().BeginCreate(ctx, "test-rg", "X00", armworkloadssapvirtualinstance.SAPVirtualInstance{
		Location: to.Ptr("northeurope"),
		Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
			Configuration: &armworkloadssapvirtualinstance.DiscoveryConfiguration{
				CentralServerVMID: to.Ptr("/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/sapq20scsvm0"),
				ConfigurationType: to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDiscovery),
			},
			Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeNonProd),
			SapProduct:  to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
		},
		Tags: map[string]*string{
			"createdby": to.Ptr("abc@microsoft.com"),
			"test":      to.Ptr("abc"),
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
	// res = armworkloadssapvirtualinstance.SAPVirtualInstancesClientCreateResponse{
	// 	SAPVirtualInstance: &armworkloadssapvirtualinstance.SAPVirtualInstance{
	// 		Name: to.Ptr("Q20"),
	// 		Type: to.Ptr("microsoft.workloads/sapvirtualinstances"),
	// 		ID: to.Ptr("/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/Q20"),
	// 		Location: to.Ptr("northeurope"),
	// 		Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
	// 			Configuration: &armworkloadssapvirtualinstance.DiscoveryConfiguration{
	// 				AppLocation: to.Ptr("westeurope"),
	// 				CentralServerVMID: to.Ptr("/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg-infra/providers/Microsoft.Compute/virtualMachines/q20ascsvm0"),
	// 				ConfigurationType: to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDiscovery),
	// 			},
	// 			Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeNonProd),
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 			ManagedResourceGroupConfiguration: &armworkloadssapvirtualinstance.ManagedRGConfiguration{
	// 				Name: to.Ptr("mrg-Q20-5b0097"),
	// 			},
	// 			ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 			SapProduct: to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
	// 			State: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStateRegistrationComplete),
	// 			Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
	// 		},
	// 		SystemData: &armworkloadssapvirtualinstance.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-27T08:22:52.3318907Z"); return t}()),
	// 			CreatedBy: to.Ptr("abc@microsoft.com"),
	// 			CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-27T11:44:17.9310503Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("abc@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
