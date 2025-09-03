package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2025-05-01-preview/Entities_CreateOrUpdate.json
func ExampleEntitiesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("4980D7D5-4E07-47AD-AD34-E76C6BC9F061", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().CreateOrUpdate(ctx, "rgopenapi", "myHealthModel", "uszrxbdkxesdrxhmagmzywebgbjj", armcloudhealth.Entity{
		Properties: &armcloudhealth.EntityProperties{
			DisplayName: to.Ptr("My entity"),
			CanvasPosition: &armcloudhealth.EntityCoordinates{
				X: to.Ptr[float32](14),
				Y: to.Ptr[float32](13),
			},
			Icon: &armcloudhealth.IconDefinition{
				IconName:   to.Ptr("Custom"),
				CustomData: to.Ptr("rcitntvapruccrhtxmkqjphbxunkz"),
			},
			HealthObjective: to.Ptr[float32](62),
			Impact:          to.Ptr(armcloudhealth.EntityImpactStandard),
			Labels: map[string]*string{
				"key1376": to.Ptr("ixfvzsfnpvkkbrce"),
			},
			Signals: &armcloudhealth.SignalGroup{
				AzureResource: &armcloudhealth.AzureResourceSignalGroup{
					SignalAssignments: []*armcloudhealth.SignalAssignment{
						{
							SignalDefinitions: []*string{
								to.Ptr("sigdef1"),
							},
						},
					},
					AuthenticationSetting: to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
					AzureResourceID:       to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1"),
				},
				AzureLogAnalytics: &armcloudhealth.LogAnalyticsSignalGroup{
					SignalAssignments: []*armcloudhealth.SignalAssignment{
						{
							SignalDefinitions: []*string{
								to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
							},
						},
					},
					AuthenticationSetting:           to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
					LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace"),
				},
				AzureMonitorWorkspace: &armcloudhealth.AzureMonitorWorkspaceSignalGroup{
					SignalAssignments: []*armcloudhealth.SignalAssignment{
						{
							SignalDefinitions: []*string{
								to.Ptr("sigdef2"),
							},
						},
						{
							SignalDefinitions: []*string{
								to.Ptr("sigdef3"),
							},
						},
					},
					AuthenticationSetting:           to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
					AzureMonitorWorkspaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace"),
				},
				Dependencies: &armcloudhealth.DependenciesSignalGroup{
					AggregationType: to.Ptr(armcloudhealth.DependenciesAggregationTypeWorstOf),
				},
			},
			Alerts: &armcloudhealth.EntityAlerts{
				Unhealthy: &armcloudhealth.AlertConfiguration{
					Severity:    to.Ptr(armcloudhealth.AlertSeveritySev1),
					Description: to.Ptr("Alert description"),
					ActionGroupIDs: []*string{
						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup"),
					},
				},
				Degraded: &armcloudhealth.AlertConfiguration{
					Severity:    to.Ptr(armcloudhealth.AlertSeveritySev4),
					Description: to.Ptr("Alert description"),
					ActionGroupIDs: []*string{
						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup"),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.EntitiesClientCreateOrUpdateResponse{
	// 	Entity: &armcloudhealth.Entity{
	// 		Properties: &armcloudhealth.EntityProperties{
	// 			ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
	// 			DisplayName: to.Ptr("My entity"),
	// 			CanvasPosition: &armcloudhealth.EntityCoordinates{
	// 				X: to.Ptr[float32](14),
	// 				Y: to.Ptr[float32](13),
	// 			},
	// 			Icon: &armcloudhealth.IconDefinition{
	// 				IconName: to.Ptr("Custom"),
	// 				CustomData: to.Ptr("rcitntvapruccrhtxmkqjphbxunkz"),
	// 			},
	// 			HealthObjective: to.Ptr[float32](62),
	// 			Impact: to.Ptr(armcloudhealth.EntityImpactStandard),
	// 			Labels: map[string]*string{
	// 				"key1376": to.Ptr("ixfvzsfnpvkkbrce"),
	// 			},
	// 			Signals: &armcloudhealth.SignalGroup{
	// 				AzureResource: &armcloudhealth.AzureResourceSignalGroup{
	// 					SignalAssignments: []*armcloudhealth.SignalAssignment{
	// 						{
	// 							SignalDefinitions: []*string{
	// 								to.Ptr("sigdef1"),
	// 							},
	// 						},
	// 					},
	// 					AuthenticationSetting: to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
	// 					AzureResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace"),
	// 				},
	// 				AzureMonitorWorkspace: &armcloudhealth.AzureMonitorWorkspaceSignalGroup{
	// 					SignalAssignments: []*armcloudhealth.SignalAssignment{
	// 						{
	// 							SignalDefinitions: []*string{
	// 								to.Ptr("sigdef2"),
	// 							},
	// 						},
	// 						{
	// 							SignalDefinitions: []*string{
	// 								to.Ptr("sigdef3"),
	// 							},
	// 						},
	// 					},
	// 					AuthenticationSetting: to.Ptr("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX"),
	// 					AzureMonitorWorkspaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace"),
	// 				},
	// 				Dependencies: &armcloudhealth.DependenciesSignalGroup{
	// 					AggregationType: to.Ptr(armcloudhealth.DependenciesAggregationTypeWorstOf),
	// 				},
	// 			},
	// 			DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:14.531Z"); return t}()),
	// 			Alerts: &armcloudhealth.EntityAlerts{
	// 				Unhealthy: &armcloudhealth.AlertConfiguration{
	// 					Severity: to.Ptr(armcloudhealth.AlertSeveritySev1),
	// 					Description: to.Ptr("Alert description"),
	// 					ActionGroupIDs: []*string{
	// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup"),
	// 					},
	// 				},
	// 				Degraded: &armcloudhealth.AlertConfiguration{
	// 					Severity: to.Ptr(armcloudhealth.AlertSeveritySev4),
	// 					Description: to.Ptr("Alert description"),
	// 					ActionGroupIDs: []*string{
	// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.CloudHealth/healthmodels/myHealthModel/entities/uszrxbdkxesdrxhmagmzywebgbjj"),
	// 		Name: to.Ptr("uszrxbdkxesdrxhmagmzywebgbjj"),
	// 		Type: to.Ptr("Microsoft.CloudHealth/healthmodels/entities"),
	// 		SystemData: &armcloudhealth.SystemData{
	// 			CreatedBy: to.Ptr("cbhzxxlvkmufetjjjwtk"),
	// 			CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.327Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("arz"),
	// 			LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.328Z"); return t}()),
	// 		},
	// 	},
	// }
}
