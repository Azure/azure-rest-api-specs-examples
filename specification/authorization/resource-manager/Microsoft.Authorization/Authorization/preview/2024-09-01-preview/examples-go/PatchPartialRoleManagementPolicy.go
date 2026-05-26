package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: 2024-09-01-preview/PatchPartialRoleManagementPolicy.json
func ExampleRoleManagementPoliciesClient_Update_patchPartialRoleManagementPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleManagementPoliciesClient().Update(ctx, "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368", "570c3619-7688-4b34-b290-2b8bb3ccab2a", armauthorization.RoleManagementPolicy{
		Properties: &armauthorization.RoleManagementPolicyProperties{
			Rules: []armauthorization.RoleManagementPolicyRuleClassification{
				&armauthorization.RoleManagementPolicyExpirationRule{
					ID:                   to.Ptr("Expiration_Admin_Eligibility"),
					IsExpirationRequired: to.Ptr(false),
					MaximumDuration:      to.Ptr("P180D"),
					RuleType:             to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("Admin"),
						Level:  to.Ptr("Eligibility"),
						Operations: []*string{
							to.Ptr("All"),
						},
					},
				},
				&armauthorization.RoleManagementPolicyNotificationRule{
					ID:                         to.Ptr("Notification_Admin_Admin_Eligibility"),
					IsDefaultRecipientsEnabled: to.Ptr(false),
					NotificationLevel:          to.Ptr(armauthorization.NotificationLevelCritical),
					NotificationRecipients: []*string{
						to.Ptr("admin_admin_eligible@test.com"),
					},
					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
					RecipientType:    to.Ptr(armauthorization.RecipientTypeAdmin),
					RuleType:         to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
					Target: &armauthorization.RoleManagementPolicyRuleTarget{
						Caller: to.Ptr("Admin"),
						Level:  to.Ptr("Eligibility"),
						Operations: []*string{
							to.Ptr("All"),
						},
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
	// res = armauthorization.RoleManagementPoliciesClientUpdateResponse{
	// 	RoleManagementPolicy: &armauthorization.RoleManagementPolicy{
	// 		Name: to.Ptr("570c3619-7688-4b34-b290-2b8bb3ccab2a"),
	// 		Type: to.Ptr("Microsoft.Authorization/RoleManagementPolicies"),
	// 		ID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleManagementPolicies/570c3619-7688-4b34-b290-2b8bb3ccab2a"),
	// 		Properties: &armauthorization.RoleManagementPolicyProperties{
	// 			EffectiveRules: []armauthorization.RoleManagementPolicyRuleClassification{
	// 				&armauthorization.RoleManagementPolicyExpirationRule{
	// 					ID: to.Ptr("Expiration_Admin_Eligibility"),
	// 					IsExpirationRequired: to.Ptr(false),
	// 					MaximumDuration: to.Ptr("P180D"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Eligibility"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Admin_Admin_Eligibility"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("admin_admin_eligible@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeAdmin),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Eligibility"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Requestor_Admin_Eligibility"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("requestor_admin_eligible@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeRequestor),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Eligibility"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Approver_Admin_Eligibility"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("approver_admin_eligible@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeApprover),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Eligibility"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyEnablementRule{
	// 					EnabledRules: []*armauthorization.EnablementRules{
	// 					},
	// 					ID: to.Ptr("Enablement_Admin_Eligibility"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Eligibility"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyExpirationRule{
	// 					ID: to.Ptr("Expiration_Admin_Assignment"),
	// 					IsExpirationRequired: to.Ptr(false),
	// 					MaximumDuration: to.Ptr("P90D"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyEnablementRule{
	// 					EnabledRules: []*armauthorization.EnablementRules{
	// 						to.Ptr(armauthorization.EnablementRulesJustification),
	// 						to.Ptr(armauthorization.EnablementRulesMultiFactorAuthentication),
	// 					},
	// 					ID: to.Ptr("Enablement_Admin_Assignment"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Admin_Admin_Assignment"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("admin_admin_member@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeAdmin),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Requestor_Admin_Assignment"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("requestor_admin_member@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeRequestor),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Approver_Admin_Assignment"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("approver_admin_member@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeApprover),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyExpirationRule{
	// 					ID: to.Ptr("Expiration_EndUser_Assignment"),
	// 					IsExpirationRequired: to.Ptr(true),
	// 					MaximumDuration: to.Ptr("PT7H"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyEnablementRule{
	// 					EnabledRules: []*armauthorization.EnablementRules{
	// 						to.Ptr(armauthorization.EnablementRulesJustification),
	// 						to.Ptr(armauthorization.EnablementRulesMultiFactorAuthentication),
	// 						to.Ptr(armauthorization.EnablementRulesTicketing),
	// 					},
	// 					ID: to.Ptr("Enablement_EndUser_Assignment"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyApprovalRule{
	// 					ID: to.Ptr("Approval_EndUser_Assignment"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule),
	// 					Setting: &armauthorization.ApprovalSettings{
	// 						ApprovalMode: to.Ptr(armauthorization.ApprovalModeSingleStage),
	// 						ApprovalStages: []*armauthorization.ApprovalStage{
	// 							{
	// 								ApprovalStageTimeOutInDays: to.Ptr[int32](1),
	// 								EscalationTimeInMinutes: to.Ptr[int32](0),
	// 								IsApproverJustificationRequired: to.Ptr(true),
	// 								IsEscalationEnabled: to.Ptr(false),
	// 								PrimaryApprovers: []*armauthorization.UserSet{
	// 									{
	// 										Description: to.Ptr("amansw_new_group"),
	// 										ID: to.Ptr("2385b0f3-5fa9-43cf-8ca4-b01dc97298cd"),
	// 										IsBackup: to.Ptr(false),
	// 										UserType: to.Ptr(armauthorization.UserTypeGroup),
	// 									},
	// 									{
	// 										Description: to.Ptr("amansw_group"),
	// 										ID: to.Ptr("2f4913c9-d15b-406a-9946-1d66a28f2690"),
	// 										IsBackup: to.Ptr(false),
	// 										UserType: to.Ptr(armauthorization.UserTypeGroup),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						IsApprovalRequired: to.Ptr(true),
	// 						IsApprovalRequiredForExtension: to.Ptr(false),
	// 						IsRequestorJustificationRequired: to.Ptr(true),
	// 					},
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyAuthenticationContextRule{
	// 					ClaimValue: to.Ptr(""),
	// 					ID: to.Ptr("AuthenticationContext_EndUser_Assignment"),
	// 					IsEnabled: to.Ptr(false),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Admin_EndUser_Assignment"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("admin_enduser_member@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeAdmin),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Requestor_EndUser_Assignment"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("requestor_enduser_member@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeRequestor),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Approver_EndUser_Assignment"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(true),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeApprover),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			IsOrganizationDefault: to.Ptr(false),
	// 			LastModifiedBy: &armauthorization.Principal{
	// 				DisplayName: to.Ptr("Admin"),
	// 			},
	// 			LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-17T16:35:27.91+00:00"); return t}()),
	// 			PolicyProperties: &armauthorization.PolicyProperties{
	// 				Scope: &armauthorization.PolicyPropertiesScope{
	// 					Type: to.Ptr("subscription"),
	// 					DisplayName: to.Ptr("Pay-As-You-Go"),
	// 					ID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368"),
	// 				},
	// 			},
	// 			Rules: []armauthorization.RoleManagementPolicyRuleClassification{
	// 				&armauthorization.RoleManagementPolicyExpirationRule{
	// 					ID: to.Ptr("Expiration_Admin_Eligibility"),
	// 					IsExpirationRequired: to.Ptr(false),
	// 					MaximumDuration: to.Ptr("P180D"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Eligibility"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Admin_Admin_Eligibility"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("admin_admin_eligible@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeAdmin),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Eligibility"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Requestor_Admin_Eligibility"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("requestor_admin_eligible@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeRequestor),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Eligibility"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Approver_Admin_Eligibility"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("approver_admin_eligible@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeApprover),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Eligibility"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyEnablementRule{
	// 					EnabledRules: []*armauthorization.EnablementRules{
	// 					},
	// 					ID: to.Ptr("Enablement_Admin_Eligibility"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Eligibility"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyExpirationRule{
	// 					ID: to.Ptr("Expiration_Admin_Assignment"),
	// 					IsExpirationRequired: to.Ptr(false),
	// 					MaximumDuration: to.Ptr("P90D"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyEnablementRule{
	// 					EnabledRules: []*armauthorization.EnablementRules{
	// 						to.Ptr(armauthorization.EnablementRulesJustification),
	// 						to.Ptr(armauthorization.EnablementRulesMultiFactorAuthentication),
	// 					},
	// 					ID: to.Ptr("Enablement_Admin_Assignment"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Admin_Admin_Assignment"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("admin_admin_member@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeAdmin),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Requestor_Admin_Assignment"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("requestor_admin_member@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeRequestor),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Approver_Admin_Assignment"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("approver_admin_member@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeApprover),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("Admin"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyExpirationRule{
	// 					ID: to.Ptr("Expiration_EndUser_Assignment"),
	// 					IsExpirationRequired: to.Ptr(true),
	// 					MaximumDuration: to.Ptr("PT7H"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyEnablementRule{
	// 					EnabledRules: []*armauthorization.EnablementRules{
	// 						to.Ptr(armauthorization.EnablementRulesJustification),
	// 						to.Ptr(armauthorization.EnablementRulesMultiFactorAuthentication),
	// 						to.Ptr(armauthorization.EnablementRulesTicketing),
	// 					},
	// 					ID: to.Ptr("Enablement_EndUser_Assignment"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyApprovalRule{
	// 					ID: to.Ptr("Approval_EndUser_Assignment"),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule),
	// 					Setting: &armauthorization.ApprovalSettings{
	// 						ApprovalMode: to.Ptr(armauthorization.ApprovalModeSingleStage),
	// 						ApprovalStages: []*armauthorization.ApprovalStage{
	// 							{
	// 								ApprovalStageTimeOutInDays: to.Ptr[int32](1),
	// 								EscalationTimeInMinutes: to.Ptr[int32](0),
	// 								IsApproverJustificationRequired: to.Ptr(true),
	// 								IsEscalationEnabled: to.Ptr(false),
	// 								PrimaryApprovers: []*armauthorization.UserSet{
	// 									{
	// 										Description: to.Ptr("amansw_new_group"),
	// 										ID: to.Ptr("2385b0f3-5fa9-43cf-8ca4-b01dc97298cd"),
	// 										IsBackup: to.Ptr(false),
	// 										UserType: to.Ptr(armauthorization.UserTypeGroup),
	// 									},
	// 									{
	// 										Description: to.Ptr("amansw_group"),
	// 										ID: to.Ptr("2f4913c9-d15b-406a-9946-1d66a28f2690"),
	// 										IsBackup: to.Ptr(false),
	// 										UserType: to.Ptr(armauthorization.UserTypeGroup),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						IsApprovalRequired: to.Ptr(true),
	// 						IsApprovalRequiredForExtension: to.Ptr(false),
	// 						IsRequestorJustificationRequired: to.Ptr(true),
	// 					},
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyAuthenticationContextRule{
	// 					ClaimValue: to.Ptr(""),
	// 					ID: to.Ptr("AuthenticationContext_EndUser_Assignment"),
	// 					IsEnabled: to.Ptr(false),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Admin_EndUser_Assignment"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("admin_enduser_member@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeAdmin),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Requestor_EndUser_Assignment"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(false),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationRecipients: []*string{
	// 						to.Ptr("requestor_enduser_member@test.com"),
	// 					},
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeRequestor),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 				&armauthorization.RoleManagementPolicyNotificationRule{
	// 					ID: to.Ptr("Notification_Approver_EndUser_Assignment"),
	// 					IsDefaultRecipientsEnabled: to.Ptr(true),
	// 					NotificationLevel: to.Ptr(armauthorization.NotificationLevelCritical),
	// 					NotificationType: to.Ptr(armauthorization.NotificationDeliveryMechanismEmail),
	// 					RecipientType: to.Ptr(armauthorization.RecipientTypeApprover),
	// 					RuleType: to.Ptr(armauthorization.RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule),
	// 					Target: &armauthorization.RoleManagementPolicyRuleTarget{
	// 						Caller: to.Ptr("EndUser"),
	// 						Level: to.Ptr("Assignment"),
	// 						Operations: []*string{
	// 							to.Ptr("All"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Scope: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368"),
	// 		},
	// 	},
	// }
}
