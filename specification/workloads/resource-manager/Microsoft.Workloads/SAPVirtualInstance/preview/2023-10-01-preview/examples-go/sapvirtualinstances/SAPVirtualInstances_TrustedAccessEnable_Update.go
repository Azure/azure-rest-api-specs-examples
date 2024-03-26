package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_TrustedAccessEnable_Update.json
func ExampleSAPVirtualInstancesClient_BeginUpdate_sapVirtualInstancesTrustedAccessEnableUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPVirtualInstancesClient().BeginUpdate(ctx, "test-rg", "X00", armworkloadssapvirtualinstance.UpdateSAPVirtualInstanceRequest{
		Identity: &armworkloadssapvirtualinstance.UserAssignedServiceIdentity{
			Type: to.Ptr(armworkloadssapvirtualinstance.ManagedServiceIdentityTypeNone),
		},
		Properties: &armworkloadssapvirtualinstance.UpdateSAPVirtualInstanceProperties{
			ManagedResourcesNetworkAccessType: to.Ptr(armworkloadssapvirtualinstance.ManagedResourcesNetworkAccessTypePrivate),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("svi1"),
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
	// 	Name: to.Ptr("X00"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00"),
	// 	SystemData: &armworkloadssapvirtualinstance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@xyz.com"),
	// 		CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@xyz.com"),
	// 		LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("svi1"),
	// 	},
	// 	Identity: &armworkloadssapvirtualinstance.UserAssignedServiceIdentity{
	// 		Type: to.Ptr(armworkloadssapvirtualinstance.ManagedServiceIdentityTypeNone),
	// 	},
	// 	Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
	// 		Configuration: &armworkloadssapvirtualinstance.DeploymentConfiguration{
	// 			ConfigurationType: to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDeployment),
	// 		},
	// 		Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeProd),
	// 		Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateUnknown),
	// 		ManagedResourceGroupConfiguration: &armworkloadssapvirtualinstance.ManagedRGConfiguration{
	// 			Name: to.Ptr("mrg-x00-6d875e77-e412-4d7d-9af4-8895278b4443"),
	// 		},
	// 		ManagedResourcesNetworkAccessType: to.Ptr(armworkloadssapvirtualinstance.ManagedResourcesNetworkAccessTypePrivate),
	// 		ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 		SapProduct: to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
	// 		State: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStateInfrastructureDeploymentPending),
	// 		Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusStarting),
	// 	},
	// }
}
