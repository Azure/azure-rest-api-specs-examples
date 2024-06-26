package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleManagementPolicyAssignmentByName.json
func ExampleRoleManagementPolicyAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleManagementPolicyAssignmentsClient().Get(ctx, "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368", "b959d571-f0b5-4042-88a7-01be6cb22db9_a1705bd2-3a8f-45a5-8683-466fcfd5cc24", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleManagementPolicyAssignment = armauthorization.RoleManagementPolicyAssignment{
	// 	Name: to.Ptr("b959d571-f0b5-4042-88a7-01be6cb22db9_a1705bd2-3a8f-45a5-8683-466fcfd5cc24"),
	// 	Type: to.Ptr("Microsoft.Authorization/RoleManagementPolicyAssignment"),
	// 	ID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleManagementPolicyAssignment/b959d571-f0b5-4042-88a7-01be6cb22db9_a1705bd2-3a8f-45a5-8683-466fcfd5cc24"),
	// 	Properties: &armauthorization.RoleManagementPolicyAssignmentProperties{
	// 		EffectiveRules: []armauthorization.RoleManagementPolicyRuleClassification{
	// 			&armauthorization.RoleManagementPolicyEnablementRule{
	// 				ID: to.Ptr("Enablement_Admin_Eligibility"),
	// 				RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
	// 				Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 					Caller: to.Ptr("Admin"),
	// 					Level: to.Ptr("Eligibility"),
	// 					Operations: []*string{
	// 						to.Ptr("All")},
	// 					},
	// 					EnabledRules: []*armauthorization.EnablementRules{
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyExpirationRule{
	// 					ID: to.Ptr("Expiration_Admin_Eligibility"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Eligibility"),
	// 						Operations: []*string{
	// 							to.Ptr("All")},
	// 						},
	// 						IsExpirationRequired: to.Ptr(true),
	// 						MaximumDuration: to.Ptr("P90D"),
	// 					},
	// 					&armauthorization.RoleManagementPolicyNotificationRule{
	// 						ID: to.Ptr("Notification_Admin_Admin_Eligibility"),
	// 						RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 						Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 							Caller: to.Ptr("Admin"),
	// 							Level: to.Ptr("Eligibility"),
	// 							Operations: []*string{
	// 								to.Ptr("All")},
	// 							},
	// 							IsDefaultRecipientsEnabled: to.Ptr(false),
	// 							NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 							NotificationRecipients: []*string{
	// 								to.Ptr("admin_admin_eligible@test.com")},
	// 								NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 								RecipientType: to.Ptr(armauthorization.RecipientTypeAdmin),
	// 							},
	// 							&armauthorization.RoleManagementPolicyNotificationRule{
	// 								ID: to.Ptr("Notification_Requestor_Admin_Eligibility"),
	// 								RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 								Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 									Caller: to.Ptr("Admin"),
	// 									Level: to.Ptr("Eligibility"),
	// 									Operations: []*string{
	// 										to.Ptr("All")},
	// 									},
	// 									IsDefaultRecipientsEnabled: to.Ptr(false),
	// 									NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 									NotificationRecipients: []*string{
	// 										to.Ptr("requestor_admin_eligible@test.com")},
	// 										NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 										RecipientType: to.Ptr(armauthorization.RecipientTypeRequestor),
	// 									},
	// 									&armauthorization.RoleManagementPolicyNotificationRule{
	// 										ID: to.Ptr("Notification_Approver_Admin_Eligibility"),
	// 										RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 										Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 											Caller: to.Ptr("Admin"),
	// 											Level: to.Ptr("Eligibility"),
	// 											Operations: []*string{
	// 												to.Ptr("All")},
	// 											},
	// 											IsDefaultRecipientsEnabled: to.Ptr(false),
	// 											NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 											NotificationRecipients: []*string{
	// 												to.Ptr("approver_admin_eligible@test.com")},
	// 												NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 												RecipientType: to.Ptr(armauthorization.RecipientTypeApprover),
	// 											},
	// 											&armauthorization.RoleManagementPolicyEnablementRule{
	// 												ID: to.Ptr("Enablement_Admin_Assignment"),
	// 												RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
	// 												Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 													Caller: to.Ptr("Admin"),
	// 													Level: to.Ptr("Assignment"),
	// 													Operations: []*string{
	// 														to.Ptr("All")},
	// 													},
	// 													EnabledRules: []*armauthorization.EnablementRules{
	// 														to.Ptr(armauthorization.EnablementRulesMultiFactorAuthentication),
	// 														to.Ptr(armauthorization.EnablementRulesJustification)},
	// 													},
	// 													&armauthorization.RoleManagementPolicyExpirationRule{
	// 														ID: to.Ptr("Expiration_Admin_Assignment"),
	// 														RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
	// 														Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 															Caller: to.Ptr("Admin"),
	// 															Level: to.Ptr("Assignment"),
	// 															Operations: []*string{
	// 																to.Ptr("All")},
	// 															},
	// 															IsExpirationRequired: to.Ptr(false),
	// 															MaximumDuration: to.Ptr("P90D"),
	// 														},
	// 														&armauthorization.RoleManagementPolicyNotificationRule{
	// 															ID: to.Ptr("Notification_Admin_Admin_Assignment"),
	// 															RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 															Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 																Caller: to.Ptr("Admin"),
	// 																Level: to.Ptr("Assignment"),
	// 																Operations: []*string{
	// 																	to.Ptr("All")},
	// 																},
	// 																IsDefaultRecipientsEnabled: to.Ptr(false),
	// 																NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 																NotificationRecipients: []*string{
	// 																	to.Ptr("admin_admin_member@test.com")},
	// 																	NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 																	RecipientType: to.Ptr(armauthorization.RecipientTypeAdmin),
	// 																},
	// 																&armauthorization.RoleManagementPolicyNotificationRule{
	// 																	ID: to.Ptr("Notification_Requestor_Admin_Assignment"),
	// 																	RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 																	Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 																		Caller: to.Ptr("Admin"),
	// 																		Level: to.Ptr("Assignment"),
	// 																		Operations: []*string{
	// 																			to.Ptr("All")},
	// 																		},
	// 																		IsDefaultRecipientsEnabled: to.Ptr(false),
	// 																		NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 																		NotificationRecipients: []*string{
	// 																			to.Ptr("requestor_admin_member@test.com")},
	// 																			NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 																			RecipientType: to.Ptr(armauthorization.RecipientTypeRequestor),
	// 																		},
	// 																		&armauthorization.RoleManagementPolicyNotificationRule{
	// 																			ID: to.Ptr("Notification_Approver_Admin_Assignment"),
	// 																			RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 																			Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 																				Caller: to.Ptr("Admin"),
	// 																				Level: to.Ptr("Assignment"),
	// 																				Operations: []*string{
	// 																					to.Ptr("All")},
	// 																				},
	// 																				IsDefaultRecipientsEnabled: to.Ptr(false),
	// 																				NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 																				NotificationRecipients: []*string{
	// 																					to.Ptr("approver_admin_member@test.com")},
	// 																					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 																					RecipientType: to.Ptr(armauthorization.RecipientTypeApprover),
	// 																				},
	// 																				&armauthorization.RoleManagementPolicyApprovalRule{
	// 																					ID: to.Ptr("Approval_EndUser_Assignment"),
	// 																					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule),
	// 																					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 																						Caller: to.Ptr("EndUser"),
	// 																						Level: to.Ptr("Assignment"),
	// 																						Operations: []*string{
	// 																							to.Ptr("All")},
	// 																						},
	// 																						Setting: &armauthorization.ApprovalSettings{
	// 																							ApprovalMode: to.Ptr(armauthorization.ApprovalModeSingleStage),
	// 																							ApprovalStages: []*armauthorization.ApprovalStage{
	// 																								{
	// 																									ApprovalStageTimeOutInDays: to.Ptr[int32](1),
	// 																									EscalationTimeInMinutes: to.Ptr[int32](0),
	// 																									IsApproverJustificationRequired: to.Ptr(true),
	// 																									IsEscalationEnabled: to.Ptr(false),
	// 																									PrimaryApprovers: []*armauthorization.UserSet{
	// 																										{
	// 																											Description: to.Ptr("amansw_new_group"),
	// 																											ID: to.Ptr("2385b0f3-5fa9-43cf-8ca4-b01dc97298cd"),
	// 																											IsBackup: to.Ptr(false),
	// 																											UserType: to.Ptr(armauthorization.UserTypeGroup),
	// 																										},
	// 																										{
	// 																											Description: to.Ptr("amansw_group"),
	// 																											ID: to.Ptr("2f4913c9-d15b-406a-9946-1d66a28f2690"),
	// 																											IsBackup: to.Ptr(false),
	// 																											UserType: to.Ptr(armauthorization.UserTypeGroup),
	// 																									}},
	// 																							}},
	// 																							IsApprovalRequired: to.Ptr(true),
	// 																							IsApprovalRequiredForExtension: to.Ptr(false),
	// 																							IsRequestorJustificationRequired: to.Ptr(true),
	// 																						},
	// 																					},
	// 																					&armauthorization.RoleManagementPolicyAuthenticationContextRule{
	// 																						ID: to.Ptr("AuthenticationContext_EndUser_Assignment"),
	// 																						RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule),
	// 																						Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 																							Caller: to.Ptr("EndUser"),
	// 																							Level: to.Ptr("Assignment"),
	// 																							Operations: []*string{
	// 																								to.Ptr("All")},
	// 																							},
	// 																							ClaimValue: to.Ptr(""),
	// 																							IsEnabled: to.Ptr(false),
	// 																						},
	// 																						&armauthorization.RoleManagementPolicyEnablementRule{
	// 																							ID: to.Ptr("Enablement_EndUser_Assignment"),
	// 																							RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
	// 																							Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 																								Caller: to.Ptr("EndUser"),
	// 																								Level: to.Ptr("Assignment"),
	// 																								Operations: []*string{
	// 																									to.Ptr("All")},
	// 																								},
	// 																								EnabledRules: []*armauthorization.EnablementRules{
	// 																									to.Ptr(armauthorization.EnablementRulesMultiFactorAuthentication),
	// 																									to.Ptr(armauthorization.EnablementRulesJustification),
	// 																									to.Ptr(armauthorization.EnablementRulesTicketing)},
	// 																								},
	// 																								&armauthorization.RoleManagementPolicyExpirationRule{
	// 																									ID: to.Ptr("Expiration_EndUser_Assignment"),
	// 																									RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
	// 																									Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 																										Caller: to.Ptr("EndUser"),
	// 																										Level: to.Ptr("Assignment"),
	// 																										Operations: []*string{
	// 																											to.Ptr("All")},
	// 																										},
	// 																										IsExpirationRequired: to.Ptr(true),
	// 																										MaximumDuration: to.Ptr("PT7H"),
	// 																									},
	// 																									&armauthorization.RoleManagementPolicyNotificationRule{
	// 																										ID: to.Ptr("Notification_Admin_EndUser_Assignment"),
	// 																										RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 																										Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 																											Caller: to.Ptr("EndUser"),
	// 																											Level: to.Ptr("Assignment"),
	// 																											Operations: []*string{
	// 																												to.Ptr("All")},
	// 																											},
	// 																											IsDefaultRecipientsEnabled: to.Ptr(false),
	// 																											NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 																											NotificationRecipients: []*string{
	// 																												to.Ptr("admin_enduser_member@test.com")},
	// 																												NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 																												RecipientType: to.Ptr(armauthorization.RecipientTypeAdmin),
	// 																											},
	// 																											&armauthorization.RoleManagementPolicyNotificationRule{
	// 																												ID: to.Ptr("Notification_Requestor_EndUser_Assignment"),
	// 																												RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 																												Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 																													Caller: to.Ptr("EndUser"),
	// 																													Level: to.Ptr("Assignment"),
	// 																													Operations: []*string{
	// 																														to.Ptr("All")},
	// 																													},
	// 																													IsDefaultRecipientsEnabled: to.Ptr(false),
	// 																													NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 																													NotificationRecipients: []*string{
	// 																														to.Ptr("requestor_enduser_member@test.com")},
	// 																														NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 																														RecipientType: to.Ptr(armauthorization.RecipientTypeRequestor),
	// 																													},
	// 																													&armauthorization.RoleManagementPolicyNotificationRule{
	// 																														ID: to.Ptr("Notification_Approver_EndUser_Assignment"),
	// 																														RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 																														Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 																															Caller: to.Ptr("EndUser"),
	// 																															Level: to.Ptr("Assignment"),
	// 																															Operations: []*string{
	// 																																to.Ptr("All")},
	// 																															},
	// 																															IsDefaultRecipientsEnabled: to.Ptr(true),
	// 																															NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 																															NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 																															RecipientType: to.Ptr(armauthorization.RecipientTypeApprover),
	// 																													}},
	// 																													PolicyAssignmentProperties: &armauthorization.PolicyAssignmentProperties{
	// 																														Policy: &armauthorization.PolicyAssignmentPropertiesPolicy{
	// 																															ID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleManagementPolicies/b959d571-f0b5-4042-88a7-01be6cb22db9"),
	// 																															LastModifiedBy: &armauthorization.Principal{
	// 																																DisplayName: to.Ptr("Admin"),
	// 																															},
	// 																														},
	// 																														RoleDefinition: &armauthorization.PolicyAssignmentPropertiesRoleDefinition{
	// 																															Type: to.Ptr("BuiltInRole"),
	// 																															DisplayName: to.Ptr("FHIR Data Converter"),
	// 																															ID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleDefinitions/a1705bd2-3a8f-45a5-8683-466fcfd5cc24"),
	// 																														},
	// 																														Scope: &armauthorization.PolicyAssignmentPropertiesScope{
	// 																															Type: to.Ptr("subscription"),
	// 																															DisplayName: to.Ptr("Pay-As-You-Go"),
	// 																															ID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368"),
	// 																														},
	// 																													},
	// 																													PolicyID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleManagementPolicies/b959d571-f0b5-4042-88a7-01be6cb22db9"),
	// 																													RoleDefinitionID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleDefinitions/a1705bd2-3a8f-45a5-8683-466fcfd5cc24"),
	// 																													Scope: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368"),
	// 																												},
	// 																											}
}
