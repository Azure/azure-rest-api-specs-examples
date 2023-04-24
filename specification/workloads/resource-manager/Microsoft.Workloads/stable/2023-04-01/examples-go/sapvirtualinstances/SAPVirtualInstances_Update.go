package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Update.json
func ExampleSAPVirtualInstancesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSAPVirtualInstancesClient().Update(ctx, "test-rg", "X00", armworkloads.UpdateSAPVirtualInstanceRequest{
		Identity: &armworkloads.UserAssignedServiceIdentity{
			Type: to.Ptr(armworkloads.ManagedServiceIdentityTypeNone),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("svi1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPVirtualInstance = armworkloads.SAPVirtualInstance{
	// 	Name: to.Ptr("X00"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00"),
	// 	SystemData: &armworkloads.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@xyz.com"),
	// 		CreatedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@xyz.com"),
	// 		LastModifiedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westcentralus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("svi1"),
	// 	},
	// 	Identity: &armworkloads.UserAssignedServiceIdentity{
	// 		Type: to.Ptr(armworkloads.ManagedServiceIdentityTypeNone),
	// 	},
	// 	Properties: &armworkloads.SAPVirtualInstanceProperties{
	// 		Configuration: &armworkloads.DeploymentConfiguration{
	// 			ConfigurationType: to.Ptr(armworkloads.SAPConfigurationTypeDeployment),
	// 		},
	// 		Environment: to.Ptr(armworkloads.SAPEnvironmentTypeProd),
	// 		Health: to.Ptr(armworkloads.SAPHealthStateUnknown),
	// 		ManagedResourceGroupConfiguration: &armworkloads.ManagedRGConfiguration{
	// 			Name: to.Ptr("mrg-x00-6d875e77-e412-4d7d-9af4-8895278b4443"),
	// 		},
	// 		ProvisioningState: to.Ptr(armworkloads.SapVirtualInstanceProvisioningStateSucceeded),
	// 		SapProduct: to.Ptr(armworkloads.SAPProductTypeS4HANA),
	// 		State: to.Ptr(armworkloads.SAPVirtualInstanceStateInfrastructureDeploymentPending),
	// 		Status: to.Ptr(armworkloads.SAPVirtualInstanceStatusStarting),
	// 	},
	// }
}
