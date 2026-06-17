package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/Drills_Update_MaximumSet_Gen.json
func ExampleDrillsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDrillsClient().BeginUpdate(ctx, "sampleServiceGroupName", "drill1", armresiliencemanagement.DrillUpdate{
		Properties: &armresiliencemanagement.DrillUpdateProperties{
			RbacSetupMode: to.Ptr(armresiliencemanagement.RBACSetupModeAutomatedCustomRole),
			RecoveryPlanProperties: &armresiliencemanagement.RecoveryPlanPropertiesOfDrill{
				Identity: &armresiliencemanagement.AssociatedIdentity{
					Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
					UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
				},
			},
			MonitoringProperties: &armresiliencemanagement.MonitoringPropertiesOfDrill{
				Identity: &armresiliencemanagement.AssociatedIdentity{
					Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
					UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
				},
			},
			ChaosResourceProperties: &armresiliencemanagement.ChaosResourcePropertiesOfDrill{
				Identity: &armresiliencemanagement.AssociatedIdentity{
					Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
					UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
				},
				ChaosResourceIdentityForFaults: &armresiliencemanagement.AssociatedIdentity{
					Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
					UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
				},
			},
			DrillAssetProperties: &armresiliencemanagement.AssetPropertiesOfDrill{
				Subscription: to.Ptr("pxlmwjuhcif"),
				Region:       to.Ptr("zuvwzxnbqyzdkthrewruw"),
			},
		},
		Identity: &armresiliencemanagement.ManagedServiceIdentity{
			Type:                   to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armresiliencemanagement.UserAssignedIdentity{},
		},
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
	// res = armresiliencemanagement.DrillsClientUpdateResponse{
	// }
}
