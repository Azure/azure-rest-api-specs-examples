package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/RecoveryPlans_CreateOrUpdate_MaximumSet_Gen.json
func ExampleRecoveryPlansClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRecoveryPlansClient().BeginCreateOrUpdate(ctx, "sampleServiceGroupName", "samplePlanName", armresiliencemanagement.RecoveryPlan{
		Properties: &armresiliencemanagement.RecoveryPlanProperties{
			PlanType:        to.Ptr(armresiliencemanagement.RecoveryPlanTypeRegional),
			PlanState:       to.Ptr(armresiliencemanagement.RecoveryPlanStateUnderEdit),
			PlanDescription: to.Ptr("Sample Plan"),
			RecoveryGroupsSetting: &armresiliencemanagement.RecoveryGroupsSetting{
				DefaultGroup: &armresiliencemanagement.RecoveryGroup{
					Properties: &armresiliencemanagement.RecoveryGroupProperties{
						GroupUniqueID: to.Ptr("b7e2a1c4-9f3b-4e2d-8c6a-2f7e4d1b5a9f"),
						OrderID:       to.Ptr[int32](0),
						Description:   to.Ptr("sample recoverygroup"),
						PreActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
								Name:             to.Ptr("sample-group-action"),
								Description:      to.Ptr("sample group action instructions"),
								Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
								TimeoutInMinutes: to.Ptr[int32](29),
								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
								Parameters: map[string]*string{
									"key7795": to.Ptr("uvapupcbbdgow"),
								},
								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
									Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
								},
							},
						},
						PostActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
								Name:             to.Ptr("sample-group-action"),
								Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
								TimeoutInMinutes: to.Ptr[int32](29),
								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
								Parameters: map[string]*string{
									"key7795": to.Ptr("uvapupcbbdgow"),
								},
								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
									Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
								},
							},
							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
								Name:             to.Ptr("sample-group-action"),
								Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
								TimeoutInMinutes: to.Ptr[int32](29),
								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
								Parameters: map[string]*string{
									"key7795": to.Ptr("uvapupcbbdgow"),
								},
								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
									Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
								},
							},
							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
								Name:             to.Ptr("sample-group-action"),
								Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
								TimeoutInMinutes: to.Ptr[int32](29),
								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
								Parameters: map[string]*string{
									"key7795": to.Ptr("uvapupcbbdgow"),
								},
								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
									Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
								},
							},
							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
								Name:             to.Ptr("sample-group-action"),
								Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
								TimeoutInMinutes: to.Ptr[int32](29),
								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
								Parameters: map[string]*string{
									"key7795": to.Ptr("uvapupcbbdgow"),
								},
								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
									Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
								},
							},
							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
								Name:             to.Ptr("sample-group-action"),
								Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
								TimeoutInMinutes: to.Ptr[int32](29),
								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
								Parameters: map[string]*string{
									"key7795": to.Ptr("uvapupcbbdgow"),
								},
								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
									Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
								},
							},
							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
								Name:             to.Ptr("sample-group-action"),
								Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
								TimeoutInMinutes: to.Ptr[int32](29),
								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
								Parameters: map[string]*string{
									"key7795": to.Ptr("uvapupcbbdgow"),
								},
								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
									Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
								},
							},
							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
								Name:             to.Ptr("sample-group-action"),
								Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
								TimeoutInMinutes: to.Ptr[int32](29),
								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
								Parameters: map[string]*string{
									"key7795": to.Ptr("uvapupcbbdgow"),
								},
								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
									Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
								},
							},
							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
								Name:             to.Ptr("sample-group-action"),
								Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
								TimeoutInMinutes: to.Ptr[int32](29),
								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
								Parameters: map[string]*string{
									"key7795": to.Ptr("uvapupcbbdgow"),
								},
								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
									Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
								},
							},
							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
								Name:             to.Ptr("sample-group-action"),
								Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
								TimeoutInMinutes: to.Ptr[int32](29),
								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
								Parameters: map[string]*string{
									"key7795": to.Ptr("uvapupcbbdgow"),
								},
								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
									Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
								},
							},
						},
					},
				},
				AdditionalGroups: []*armresiliencemanagement.RecoveryGroup{
					{
						Properties: &armresiliencemanagement.RecoveryGroupProperties{
							GroupUniqueID: to.Ptr("b7e2a1c4-9f3b-4e2d-8c6a-2f7e4d1b5a9f"),
							OrderID:       to.Ptr[int32](1),
							Description:   to.Ptr("sample recoverygroup"),
							PreActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
							},
							PostActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
							},
						},
					},
					{
						Properties: &armresiliencemanagement.RecoveryGroupProperties{
							GroupUniqueID: to.Ptr("b7e2a1c4-9f3b-4e2d-8c6a-2f7e4d1b5a9f"),
							OrderID:       to.Ptr[int32](1),
							Description:   to.Ptr("sample recoverygroup"),
							PreActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
							},
							PostActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
							},
						},
					},
					{
						Properties: &armresiliencemanagement.RecoveryGroupProperties{
							GroupUniqueID: to.Ptr("b7e2a1c4-9f3b-4e2d-8c6a-2f7e4d1b5a9f"),
							OrderID:       to.Ptr[int32](1),
							Description:   to.Ptr("sample recoverygroup"),
							PreActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
							},
							PostActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
									Name:             to.Ptr("sample-group-action"),
									Type:             to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
									TimeoutInMinutes: to.Ptr[int32](29),
									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
									Parameters: map[string]*string{
										"key7795": to.Ptr("uvapupcbbdgow"),
									},
									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
										Type:                 to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
									},
								},
							},
						},
					},
				},
			},
		},
		Identity: &armresiliencemanagement.ManagedServiceIdentity{
			Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armresiliencemanagement.UserAssignedIdentity{
				"/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1": {},
			},
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
	// res = armresiliencemanagement.RecoveryPlansClientCreateOrUpdateResponse{
	// 	RecoveryPlan: armresiliencemanagement.RecoveryPlan{
	// 		Properties: &armresiliencemanagement.RecoveryPlanProperties{
	// 			ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
	// 			PlanType: to.Ptr(armresiliencemanagement.RecoveryPlanTypeRegional),
	// 			PlanState: to.Ptr(armresiliencemanagement.RecoveryPlanStateUnderEdit),
	// 			PlanDescription: to.Ptr("etxqslinfpklqndvbpyegpqv"),
	// 			RecoveryGroupsSetting: &armresiliencemanagement.RecoveryGroupsSetting{
	// 				DefaultGroup: &armresiliencemanagement.RecoveryGroup{
	// 					Properties: &armresiliencemanagement.RecoveryGroupProperties{
	// 						GroupUniqueID: to.Ptr("b7e2a1c4-9f3b-4e2d-8c6a-2f7e4d1b5a9f"),
	// 						OrderID: to.Ptr[int32](0),
	// 						Description: to.Ptr("sample recoverygroup"),
	// 						PreActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
	// 							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 								Name: to.Ptr("sample-group-action"),
	// 								Description: to.Ptr("sample group action instructions"),
	// 								Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 								TimeoutInMinutes: to.Ptr[int32](29),
	// 								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 								Parameters: map[string]*string{
	// 									"key7795": to.Ptr("uvapupcbbdgow"),
	// 								},
	// 								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 									Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 								},
	// 							},
	// 						},
	// 						PostActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
	// 							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 								Name: to.Ptr("sample-group-action"),
	// 								Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 								TimeoutInMinutes: to.Ptr[int32](29),
	// 								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 								Parameters: map[string]*string{
	// 									"key7795": to.Ptr("uvapupcbbdgow"),
	// 								},
	// 								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 									Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 								},
	// 							},
	// 							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 								Name: to.Ptr("sample-group-action"),
	// 								Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 								TimeoutInMinutes: to.Ptr[int32](29),
	// 								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 								Parameters: map[string]*string{
	// 									"key7795": to.Ptr("uvapupcbbdgow"),
	// 								},
	// 								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 									Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 								},
	// 							},
	// 							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 								Name: to.Ptr("sample-group-action"),
	// 								Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 								TimeoutInMinutes: to.Ptr[int32](29),
	// 								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 								Parameters: map[string]*string{
	// 									"key7795": to.Ptr("uvapupcbbdgow"),
	// 								},
	// 								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 									Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 								},
	// 							},
	// 							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 								Name: to.Ptr("sample-group-action"),
	// 								Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 								TimeoutInMinutes: to.Ptr[int32](29),
	// 								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 								Parameters: map[string]*string{
	// 									"key7795": to.Ptr("uvapupcbbdgow"),
	// 								},
	// 								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 									Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 								},
	// 							},
	// 							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 								Name: to.Ptr("sample-group-action"),
	// 								Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 								TimeoutInMinutes: to.Ptr[int32](29),
	// 								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 								Parameters: map[string]*string{
	// 									"key7795": to.Ptr("uvapupcbbdgow"),
	// 								},
	// 								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 									Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 								},
	// 							},
	// 							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 								Name: to.Ptr("sample-group-action"),
	// 								Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 								TimeoutInMinutes: to.Ptr[int32](29),
	// 								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 								Parameters: map[string]*string{
	// 									"key7795": to.Ptr("uvapupcbbdgow"),
	// 								},
	// 								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 									Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 								},
	// 							},
	// 							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 								Name: to.Ptr("sample-group-action"),
	// 								Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 								TimeoutInMinutes: to.Ptr[int32](29),
	// 								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 								Parameters: map[string]*string{
	// 									"key7795": to.Ptr("uvapupcbbdgow"),
	// 								},
	// 								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 									Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 								},
	// 							},
	// 							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 								Name: to.Ptr("sample-group-action"),
	// 								Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 								TimeoutInMinutes: to.Ptr[int32](29),
	// 								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 								Parameters: map[string]*string{
	// 									"key7795": to.Ptr("uvapupcbbdgow"),
	// 								},
	// 								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 									Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 								},
	// 							},
	// 							&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 								Name: to.Ptr("sample-group-action"),
	// 								Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 								TimeoutInMinutes: to.Ptr[int32](29),
	// 								ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 								Parameters: map[string]*string{
	// 									"key7795": to.Ptr("uvapupcbbdgow"),
	// 								},
	// 								AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 									Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 									UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 					ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName/recoveryGroups/12345678-1234-1234-1234-123456789012"),
	// 					Name: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 					Type: to.Ptr("Microsoft.AzureResilienceManagement/recoveryPlans/recoveryGroups"),
	// 					SystemData: &armresiliencemanagement.SystemData{
	// 						CreatedBy: to.Ptr("wmfonl"),
	// 						CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:54.175Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("paiugykk"),
	// 						LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:54.176Z"); return t}()),
	// 					},
	// 				},
	// 				AdditionalGroups: []*armresiliencemanagement.RecoveryGroup{
	// 					{
	// 						Properties: &armresiliencemanagement.RecoveryGroupProperties{
	// 							GroupUniqueID: to.Ptr("b7e2a1c4-9f3b-4e2d-8c6a-2f7e4d1b5a9f"),
	// 							OrderID: to.Ptr[int32](1),
	// 							Description: to.Ptr("sample recoverygroup"),
	// 							PreActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 							},
	// 							PostActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName/recoveryGroups/12345678-1234-1234-1234-123456789012"),
	// 						Name: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 						Type: to.Ptr("Microsoft.AzureResilienceManagement/recoveryPlans/recoveryGroups"),
	// 						SystemData: &armresiliencemanagement.SystemData{
	// 							CreatedBy: to.Ptr("wmfonl"),
	// 							CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 							CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:54.175Z"); return t}()),
	// 							LastModifiedBy: to.Ptr("paiugykk"),
	// 							LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 							LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:54.176Z"); return t}()),
	// 						},
	// 					},
	// 					{
	// 						Properties: &armresiliencemanagement.RecoveryGroupProperties{
	// 							GroupUniqueID: to.Ptr("b7e2a1c4-9f3b-4e2d-8c6a-2f7e4d1b5a9f"),
	// 							OrderID: to.Ptr[int32](1),
	// 							Description: to.Ptr("sample recoverygroup"),
	// 							PreActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 							},
	// 							PostActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName/recoveryGroups/12345678-1234-1234-1234-123456789012"),
	// 						Name: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 						Type: to.Ptr("Microsoft.AzureResilienceManagement/recoveryPlans/recoveryGroups"),
	// 						SystemData: &armresiliencemanagement.SystemData{
	// 							CreatedBy: to.Ptr("wmfonl"),
	// 							CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 							CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:54.175Z"); return t}()),
	// 							LastModifiedBy: to.Ptr("paiugykk"),
	// 							LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 							LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:54.176Z"); return t}()),
	// 						},
	// 					},
	// 					{
	// 						Properties: &armresiliencemanagement.RecoveryGroupProperties{
	// 							GroupUniqueID: to.Ptr("b7e2a1c4-9f3b-4e2d-8c6a-2f7e4d1b5a9f"),
	// 							OrderID: to.Ptr[int32](1),
	// 							Description: to.Ptr("sample recoverygroup"),
	// 							PreActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 							},
	// 							PostActions: []armresiliencemanagement.RecoveryGroupBaseActionClassification{
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 								&armresiliencemanagement.RecoveryGroupCustomRunbookAction{
	// 									Name: to.Ptr("sample-group-action"),
	// 									Type: to.Ptr(armresiliencemanagement.RecoveryGroupActionTypeCustomRunbook),
	// 									TimeoutInMinutes: to.Ptr[int32](29),
	// 									ActionResourceID: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.Automation/automationAccounts/sampleAccount/runbooks/sameplRunbooks1"),
	// 									Parameters: map[string]*string{
	// 										"key7795": to.Ptr("uvapupcbbdgow"),
	// 									},
	// 									AssociatedIdentity: &armresiliencemanagement.AssociatedIdentity{
	// 										Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 										UserAssignedIdentity: to.Ptr("/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName/recoveryGroups/12345678-1234-1234-1234-123456789012"),
	// 						Name: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 						Type: to.Ptr("Microsoft.AzureResilienceManagement/recoveryPlans/recoveryGroups"),
	// 						SystemData: &armresiliencemanagement.SystemData{
	// 							CreatedBy: to.Ptr("wmfonl"),
	// 							CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 							CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:54.175Z"); return t}()),
	// 							LastModifiedBy: to.Ptr("paiugykk"),
	// 							LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 							LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:54.176Z"); return t}()),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			LatestFailoverStatus: &armresiliencemanagement.RecoveryPlanFailoverOperationStatus{
	// 				LastExecutedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:56.306Z"); return t}()),
	// 				OperationStatus: to.Ptr(armresiliencemanagement.RecoveryOperationStatus("FailoverFailed")),
	// 				ErrorDetails: &armresiliencemanagement.ErrorDetail{
	// 					Code: to.Ptr("ResourceInNotProtectedState"),
	// 					Message: to.Ptr("Resource in not protected with any recovery solution."),
	// 					Target: to.Ptr("Please make sure resource is protected with any recovery solution."),
	// 					Details: []*armresiliencemanagement.ErrorDetail{
	// 					},
	// 					AdditionalInfo: []*armresiliencemanagement.ErrorAdditionalInfo{
	// 						{
	// 							Type: to.Ptr("Unable to detect resource protections with any recovery solution."),
	// 							Info: map[string]any{
	// 							},
	// 						},
	// 					},
	// 				},
	// 				RecoveryTimeActual: to.Ptr("PT30M"),
	// 			},
	// 			LatestValidationStatus: &armresiliencemanagement.RecoveryPlanOperationStatus{
	// 				LastExecutedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:56.307Z"); return t}()),
	// 				OperationStatus: to.Ptr(armresiliencemanagement.RecoveryOperationStatusValidationFailed),
	// 				ErrorDetails: &armresiliencemanagement.ErrorDetail{
	// 					Code: to.Ptr("ResourceInNotProtectedState"),
	// 					Message: to.Ptr("Resource in not protected with any recovery solution."),
	// 					Target: to.Ptr("Please make sure resource is protected with any recovery solution."),
	// 					Details: []*armresiliencemanagement.ErrorDetail{
	// 					},
	// 					AdditionalInfo: []*armresiliencemanagement.ErrorAdditionalInfo{
	// 						{
	// 							Type: to.Ptr("Unable to detect resource protections with any recovery solution."),
	// 							Info: map[string]any{
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Identity: &armresiliencemanagement.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("11111111-1111-1111-1111-000000000012"),
	// 			TenantID: to.Ptr("11111111-1111-1111-1111-000000000011"),
	// 			Type: to.Ptr(armresiliencemanagement.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armresiliencemanagement.UserAssignedIdentity{
	// 				"/subscriptions/4e88bed3-114f-443d-9975-28f64122ec5e/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1": &armresiliencemanagement.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("11111111-1111-1111-1111-000000000012"),
	// 					ClientID: to.Ptr("11111111-1111-1111-1111-000000000011"),
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName"),
	// 		Name: to.Ptr("samplePlanName"),
	// 		Type: to.Ptr("Microsoft.AzureResilienceManagement/recoveryPlans"),
	// 		SystemData: &armresiliencemanagement.SystemData{
	// 			CreatedBy: to.Ptr("wmfonl"),
	// 			CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:54.175Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("paiugykk"),
	// 			LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-15T09:19:54.176Z"); return t}()),
	// 		},
	// 	},
	// }
}
