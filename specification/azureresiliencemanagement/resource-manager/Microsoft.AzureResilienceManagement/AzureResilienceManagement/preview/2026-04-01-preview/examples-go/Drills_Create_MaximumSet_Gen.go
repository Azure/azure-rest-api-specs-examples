package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/Drills_Create_MaximumSet_Gen.json
func ExampleDrillsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDrillsClient().BeginCreate(ctx, "sampleServiceGroupName", "drill1", armresiliencemanagement.Drill{
		Properties: &armresiliencemanagement.DrillProperties{
			ExecutionState: to.Ptr(armresiliencemanagement.ExecutionStateNotRunning),
			RbacSetupMode:  to.Ptr(armresiliencemanagement.RBACSetupModeAutomatedCustomRole),
			AttentionReason: &armresiliencemanagement.AttentionReason{
				RoReadiness:             to.Ptr(armresiliencemanagement.RecoveryPlanStateUnderEdit),
				DrillUserMsi:            to.Ptr(armresiliencemanagement.ExtensionObjectStateExists),
				DrillRbacOnRecoveryPlan: to.Ptr(armresiliencemanagement.RBACStateSet),
				RbacNeededForDrillOnRecoveryPlan: []*string{
					to.Ptr("ayfyepziwdyxuwuexlamaadey"),
				},
				RecoveryPlanAndDrillResourcesState: to.Ptr(armresiliencemanagement.RelativeResourceCompositionStateInSync),
				ServiceGroupAndDrillResourcesState: to.Ptr(armresiliencemanagement.RelativeResourceCompositionStateInSync),
				RunbookFaultRbacOnTargets:          to.Ptr(armresiliencemanagement.RBACStateSet),
				IncludedResourceInDrill:            to.Ptr(armresiliencemanagement.ExtensionObjectStateExists),
				DrillRbacOnMonitoringResources:     to.Ptr(armresiliencemanagement.RBACStateSet),
				DrillMonitoringErrors: []*armresiliencemanagement.ErrorDetails{
					{
						Code:    to.Ptr("14123903"),
						Message: to.Ptr("Unable to assign Monitoring RBAC on target resource."),
					},
				},
				DrillMonitoringResources:       to.Ptr(armresiliencemanagement.ExtensionObjectStateExists),
				MonitoringRbacOnDrillResources: to.Ptr(armresiliencemanagement.RBACStateSet),
				RbacNeededForDrillOnDrillMonitoringResources: []*string{
					to.Ptr("lyffvljvuhwvxcuzyzlyo"),
				},
				RbacNeededForDrillOnDrillResources: []*string{
					to.Ptr("sajsgcweakvzfunxfzzxe"),
				},
				MissingRequiredResourceProviders: []*string{
					to.Ptr("Microsoft.Chaos"),
					to.Ptr("Microsoft.Automation"),
				},
				DrillRbacOnChaosResource: to.Ptr(armresiliencemanagement.RBACStateSet),
				RbacNeededForDrillOnChaosResource: []*string{
					to.Ptr("zabszxqjflfjgifyrtttvdpipw"),
				},
				ChaosResource:         to.Ptr(armresiliencemanagement.ExtensionObjectStateExists),
				ChaosResourceUserMsi:  to.Ptr(armresiliencemanagement.ExtensionObjectStateExists),
				RbacOnTargetResources: to.Ptr(armresiliencemanagement.RBACStateSet),
				ChaosResourceCreationFailureReasons: []*string{
					to.Ptr("cjqkzeqcktfqcpmdwoloqb"),
				},
			},
			SystemMetadata: &armresiliencemanagement.SystemMetadata{
				InitialConfig: to.Ptr(armresiliencemanagement.InitialConfigPending),
			},
			DrillType: to.Ptr(armresiliencemanagement.DrillType("DrillProperties")),
			LastRunProperties: &armresiliencemanagement.LastRunProperties{
				LastRunState:       to.Ptr(armresiliencemanagement.JobStatusNotStarted),
				LastRunAttestation: to.Ptr(armresiliencemanagement.DrillAttestationAttestedSuccess),
			},
			ExecutionReadinessState:        to.Ptr(armresiliencemanagement.ExecutionReadinessStateReady),
			ManagedOnBehalfOfConfiguration: &armresiliencemanagement.ManagedOnBehalfOfConfiguration{},
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
			DrillAssetProperties: &armresiliencemanagement.AssetPropertiesOfDrill{
				Subscription:  to.Ptr("4e88bed3-114f-443d-9975-28f64122ec5e"),
				Region:        to.Ptr("eastus"),
				ResourceGroup: to.Ptr("customDrillResourceGroup"),
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
	// res = armresiliencemanagement.DrillsClientCreateResponse{
	// 	Drill: armresiliencemanagement.Drill{
	// 		Properties: &armresiliencemanagement.DrillProperties{
	// 			ExecutionState: to.Ptr(armresiliencemanagement.ExecutionStateNotRunning),
	// 			RbacSetupMode: to.Ptr(armresiliencemanagement.RBACSetupModeAutomatedCustomRole),
	// 			AttentionReason: &armresiliencemanagement.AttentionReason{
	// 				RoReadiness: to.Ptr(armresiliencemanagement.RecoveryPlanStateUnderEdit),
	// 				DrillUserMsi: to.Ptr(armresiliencemanagement.ExtensionObjectStateExists),
	// 				DrillRbacOnRecoveryPlan: to.Ptr(armresiliencemanagement.RBACStateSet),
	// 				RbacNeededForDrillOnRecoveryPlan: []*string{
	// 					to.Ptr("ayfyepziwdyxuwuexlamaadey"),
	// 				},
	// 				RecoveryPlanAndDrillResourcesState: to.Ptr(armresiliencemanagement.RelativeResourceCompositionStateInSync),
	// 				ServiceGroupAndDrillResourcesState: to.Ptr(armresiliencemanagement.RelativeResourceCompositionStateInSync),
	// 				RunbookFaultRbacOnTargets: to.Ptr(armresiliencemanagement.RBACStateSet),
	// 				IncludedResourceInDrill: to.Ptr(armresiliencemanagement.ExtensionObjectStateExists),
	// 				DrillRbacOnMonitoringResources: to.Ptr(armresiliencemanagement.RBACStateSet),
	// 				DrillMonitoringErrors: []*armresiliencemanagement.ErrorDetails{
	// 					{
	// 						Code: to.Ptr("14123903"),
	// 						Message: to.Ptr("Unable to assign Monitoring RBAC on target resource."),
	// 					},
	// 				},
	// 				DrillMonitoringResources: to.Ptr(armresiliencemanagement.ExtensionObjectStateExists),
	// 				MonitoringRbacOnDrillResources: to.Ptr(armresiliencemanagement.RBACStateSet),
	// 				RbacNeededForDrillOnDrillMonitoringResources: []*string{
	// 					to.Ptr("lyffvljvuhwvxcuzyzlyo"),
	// 				},
	// 				RbacNeededForDrillOnDrillResources: []*string{
	// 					to.Ptr("sajsgcweakvzfunxfzzxe"),
	// 				},
	// 				MissingRequiredResourceProviders: []*string{
	// 					to.Ptr("Microsoft.Chaos"),
	// 					to.Ptr("Microsoft.Automation"),
	// 				},
	// 				DrillRbacOnChaosResource: to.Ptr(armresiliencemanagement.RBACStateSet),
	// 				RbacNeededForDrillOnChaosResource: []*string{
	// 					to.Ptr("zabszxqjflfjgifyrtttvdpipw"),
	// 				},
	// 				ChaosResource: to.Ptr(armresiliencemanagement.ExtensionObjectStateExists),
	// 				ChaosResourceUserMsi: to.Ptr(armresiliencemanagement.ExtensionObjectStateExists),
	// 				RbacOnTargetResources: to.Ptr(armresiliencemanagement.RBACStateSet),
	// 				ChaosResourceCreationFailureReasons: []*string{
	// 					to.Ptr("cjqkzeqcktfqcpmdwoloqb"),
	// 				},
	// 			},
	// 			SystemMetadata: &armresiliencemanagement.SystemMetadata{
	// 				InitialConfig: to.Ptr(armresiliencemanagement.InitialConfigPending),
	// 				ResourceTypeCategories: []*armresiliencemanagement.ResourceTypeCategories{
	// 					to.Ptr(armresiliencemanagement.ResourceTypeCategoriesAzureSiteRecoveryVMsPresent),
	// 				},
	// 			},
	// 			DrillType: to.Ptr(armresiliencemanagement.DrillType("DrillProperties")),
	// 			LastRunProperties: &armresiliencemanagement.LastRunProperties{
	// 				LastRunState: to.Ptr(armresiliencemanagement.JobStatusNotStarted),
	// 				LastRunAttestation: to.Ptr(armresiliencemanagement.DrillAttestationAttestedSuccess),
	// 				LastRunTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-05T05:36:37.151Z"); return t}()),
	// 				LastRunDuration: to.Ptr("PT6M"),
	// 			},
	// 			ExecutionReadinessState: to.Ptr(armresiliencemanagement.ExecutionReadinessStateReady),
	// 			ManagedOnBehalfOfConfiguration: &armresiliencemanagement.ManagedOnBehalfOfConfiguration{
	// 				MoboBrokerResources: []*armresiliencemanagement.MoboBrokerResource{
	// 					{
	// 						ID: to.Ptr("lvukttgzvssiupnypauorkyzvzf"),
	// 					},
	// 				},
	// 			},
	// 			RecoveryPlanProperties: &armresiliencemanagement.RecoveryPlanPropertiesOfDrill{
	// 				Identity: &armresiliencemanagement.AssociatedIdentity{
	// 					Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 					UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 				},
	// 				RecoveryPlanID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/rePlan1"),
	// 				RecoveryPlanResourceExcludedCount: to.Ptr[int32](29),
	// 			},
	// 			MonitoringProperties: &armresiliencemanagement.MonitoringPropertiesOfDrill{
	// 				Identity: &armresiliencemanagement.AssociatedIdentity{
	// 					Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 					UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 				},
	// 				LogAnalyticsWorkspaceID: to.Ptr("/subscriptions/191973cd-9c54-41e0-ac19-25dd9a92d5a8/resourceGroups/AzureResilienceManagementDrills/providers/microsoft.operationalinsights/workspaces/default-law-cid"),
	// 				RawMetricsDataCollectionRuleID: to.Ptr("/subscriptions/2427679b-6638-48e5-8774-6096cd849451/resourceGroups/AzureResilienceManagementDrills/providers/Microsoft.Insights/dataCollectionRules/DataCollectionRuleHealthName"),
	// 				ServiceGroupMetricsDataCollectionRuleID: to.Ptr("/subscriptions/2427679b-6638-48e5-8774-6096cd849451/resourceGroups/AzureResilienceManagementDrills/providers/Microsoft.Insights/dataCollectionRules/DataCollectionRawRuleHealthName"),
	// 				DataCollectionEndpointID: to.Ptr("/subscriptions/191973cd-9c54-41e0-ac19-25dd9a92d5a8/resourceGroups/AzureResilienceManagementDrills/providers/Microsoft.Insights/dataCollectionEndpoints/DrillDataCollectionEndpoint"),
	// 			},
	// 			DrillAssetProperties: &armresiliencemanagement.AssetPropertiesOfDrill{
	// 				Subscription: to.Ptr("4e88bed3-114f-443d-9975-28f64122ec5e"),
	// 				Region: to.Ptr("eastus"),
	// 				ResourceGroup: to.Ptr("customDrillResourceGroup"),
	// 			},
	// 			ChaosResourceProperties: &armresiliencemanagement.ChaosResourcePropertiesOfDrill{
	// 				Identity: &armresiliencemanagement.AssociatedIdentity{
	// 					Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 					UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 				},
	// 				ChaosResourceIdentityForFaults: &armresiliencemanagement.AssociatedIdentity{
	// 					Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 					UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 				},
	// 				ChaosResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourceGroups/resourceGroup1/providers/Microsoft.Chaos/workspaces/asdas"),
	// 				FaultDurationInMin: to.Ptr[int32](0),
	// 			},
	// 			ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
	// 			ServiceGroupID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName"),
	// 			LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-02-02T10:45:16.573Z"); return t}()),
	// 			LastResyncReadinessCheckTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-02-02T10:29:02.037Z"); return t}()),
	// 		},
	// 		Identity: &armresiliencemanagement.ManagedServiceIdentity{
	// 			Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armresiliencemanagement.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("b958c8d6-705d-4581-8cee-5120b43b0b16"),
	// 			TenantID: to.Ptr("0ddf1a8d-cd80-412e-b414-cdc9de46c7bf"),
	// 		},
	// 		ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/drills/drill1"),
	// 		Name: to.Ptr("drill1"),
	// 		Type: to.Ptr("Microsoft.AzureResilienceManagement/drills"),
	// 		SystemData: &armresiliencemanagement.SystemData{
	// 			CreatedBy: to.Ptr("dvnfxbuyqhvivfjddjccdtlwajfht"),
	// 			CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-06T15:03:42.796Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("lndhhaimomorael"),
	// 			LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-06T15:03:42.797Z"); return t}()),
	// 		},
	// 	},
	// }
}
