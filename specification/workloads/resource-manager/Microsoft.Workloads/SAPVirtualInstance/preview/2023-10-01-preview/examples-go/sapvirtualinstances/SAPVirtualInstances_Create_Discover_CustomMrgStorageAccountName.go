package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Create_Discover_CustomMrgStorageAccountName.json
func ExampleSAPVirtualInstancesClient_BeginCreate_registerExistingSapSystemAsVirtualInstanceForSapSolutionsWithOptionalCustomizations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPVirtualInstancesClient().BeginCreate(ctx, "test-rg", "X00", armworkloadssapvirtualinstance.SAPVirtualInstance{
		Location: to.Ptr("northeurope"),
		Tags: map[string]*string{
			"createdby": to.Ptr("abc@microsoft.com"),
			"test":      to.Ptr("abc"),
		},
		Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
			Configuration: &armworkloadssapvirtualinstance.DiscoveryConfiguration{
				ConfigurationType:           to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDiscovery),
				CentralServerVMID:           to.Ptr("/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/sapq20scsvm0"),
				ManagedRgStorageAccountName: to.Ptr("q20saacssgrs"),
			},
			Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeNonProd),
			SapProduct:  to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
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
	// res.SAPVirtualInstance = armworkloadssapvirtualinstance.SAPVirtualInstance{
	// 	Name: to.Ptr("Q20"),
	// 	Type: to.Ptr("microsoft.workloads/sapvirtualinstances"),
	// 	ID: to.Ptr("/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/Q20"),
	// 	SystemData: &armworkloadssapvirtualinstance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-27T08:22:52.331Z"); return t}()),
	// 		CreatedBy: to.Ptr("abc@microsoft.com"),
	// 		CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-27T11:44:17.931Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("abc@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("northeurope"),
	// 	Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
	// 		Configuration: &armworkloadssapvirtualinstance.DiscoveryConfiguration{
	// 			ConfigurationType: to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDiscovery),
	// 			AppLocation: to.Ptr("westeurope"),
	// 			CentralServerVMID: to.Ptr("/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg-infra/providers/Microsoft.Compute/virtualMachines/q20ascsvm0"),
	// 			ManagedRgStorageAccountName: to.Ptr("q20saacssgrs"),
	// 		},
	// 		Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeNonProd),
	// 		Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
	// 		ManagedResourceGroupConfiguration: &armworkloadssapvirtualinstance.ManagedRGConfiguration{
	// 			Name: to.Ptr("mrg-Q20-5b0097"),
	// 		},
	// 		ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 		SapProduct: to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
	// 		State: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStateRegistrationComplete),
	// 		Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
	// 	},
	// }
}
