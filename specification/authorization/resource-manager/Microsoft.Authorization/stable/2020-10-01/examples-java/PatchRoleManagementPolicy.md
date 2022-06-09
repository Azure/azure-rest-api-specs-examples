```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.authorization.fluent.models.RoleManagementPolicyInner;
import com.azure.resourcemanager.authorization.models.ApprovalMode;
import com.azure.resourcemanager.authorization.models.ApprovalSettings;
import com.azure.resourcemanager.authorization.models.ApprovalStage;
import com.azure.resourcemanager.authorization.models.EnablementRules;
import com.azure.resourcemanager.authorization.models.NotificationDeliveryMechanism;
import com.azure.resourcemanager.authorization.models.NotificationLevel;
import com.azure.resourcemanager.authorization.models.RecipientType;
import com.azure.resourcemanager.authorization.models.RoleManagementPolicyApprovalRule;
import com.azure.resourcemanager.authorization.models.RoleManagementPolicyAuthenticationContextRule;
import com.azure.resourcemanager.authorization.models.RoleManagementPolicyEnablementRule;
import com.azure.resourcemanager.authorization.models.RoleManagementPolicyExpirationRule;
import com.azure.resourcemanager.authorization.models.RoleManagementPolicyNotificationRule;
import com.azure.resourcemanager.authorization.models.RoleManagementPolicyRuleTarget;
import com.azure.resourcemanager.authorization.models.UserSet;
import com.azure.resourcemanager.authorization.models.UserType;
import java.util.Arrays;

/** Samples for RoleManagementPolicies Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/PatchRoleManagementPolicy.json
     */
    /**
     * Sample code: PatchRoleManagementPolicy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchRoleManagementPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleManagementPolicies()
            .updateWithResponse(
                "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368",
                "570c3619-7688-4b34-b290-2b8bb3ccab2a",
                new RoleManagementPolicyInner()
                    .withRules(
                        Arrays
                            .asList(
                                new RoleManagementPolicyExpirationRule()
                                    .withId("Expiration_Admin_Eligibility")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("Admin")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Eligibility"))
                                    .withIsExpirationRequired(false)
                                    .withMaximumDuration("P180D"),
                                new RoleManagementPolicyNotificationRule()
                                    .withId("Notification_Admin_Admin_Eligibility")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("Admin")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Eligibility"))
                                    .withNotificationType(NotificationDeliveryMechanism.EMAIL)
                                    .withNotificationLevel(NotificationLevel.CRITICAL)
                                    .withRecipientType(RecipientType.ADMIN)
                                    .withNotificationRecipients(Arrays.asList("admin_admin_eligible@test.com"))
                                    .withIsDefaultRecipientsEnabled(false),
                                new RoleManagementPolicyNotificationRule()
                                    .withId("Notification_Requestor_Admin_Eligibility")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("Admin")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Eligibility"))
                                    .withNotificationType(NotificationDeliveryMechanism.EMAIL)
                                    .withNotificationLevel(NotificationLevel.CRITICAL)
                                    .withRecipientType(RecipientType.REQUESTOR)
                                    .withNotificationRecipients(Arrays.asList("requestor_admin_eligible@test.com"))
                                    .withIsDefaultRecipientsEnabled(false),
                                new RoleManagementPolicyNotificationRule()
                                    .withId("Notification_Approver_Admin_Eligibility")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("Admin")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Eligibility"))
                                    .withNotificationType(NotificationDeliveryMechanism.EMAIL)
                                    .withNotificationLevel(NotificationLevel.CRITICAL)
                                    .withRecipientType(RecipientType.APPROVER)
                                    .withNotificationRecipients(Arrays.asList("approver_admin_eligible@test.com"))
                                    .withIsDefaultRecipientsEnabled(false),
                                new RoleManagementPolicyEnablementRule()
                                    .withId("Enablement_Admin_Eligibility")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("Admin")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Eligibility"))
                                    .withEnabledRules(Arrays.asList()),
                                new RoleManagementPolicyExpirationRule()
                                    .withId("Expiration_Admin_Assignment")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("Admin")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Assignment"))
                                    .withIsExpirationRequired(false)
                                    .withMaximumDuration("P90D"),
                                new RoleManagementPolicyEnablementRule()
                                    .withId("Enablement_Admin_Assignment")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("Admin")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Assignment"))
                                    .withEnabledRules(
                                        Arrays
                                            .asList(
                                                EnablementRules.JUSTIFICATION,
                                                EnablementRules.MULTI_FACTOR_AUTHENTICATION)),
                                new RoleManagementPolicyNotificationRule()
                                    .withId("Notification_Admin_Admin_Assignment")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("Admin")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Assignment"))
                                    .withNotificationType(NotificationDeliveryMechanism.EMAIL)
                                    .withNotificationLevel(NotificationLevel.CRITICAL)
                                    .withRecipientType(RecipientType.ADMIN)
                                    .withNotificationRecipients(Arrays.asList("admin_admin_member@test.com"))
                                    .withIsDefaultRecipientsEnabled(false),
                                new RoleManagementPolicyNotificationRule()
                                    .withId("Notification_Requestor_Admin_Assignment")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("Admin")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Assignment"))
                                    .withNotificationType(NotificationDeliveryMechanism.EMAIL)
                                    .withNotificationLevel(NotificationLevel.CRITICAL)
                                    .withRecipientType(RecipientType.REQUESTOR)
                                    .withNotificationRecipients(Arrays.asList("requestor_admin_member@test.com"))
                                    .withIsDefaultRecipientsEnabled(false),
                                new RoleManagementPolicyNotificationRule()
                                    .withId("Notification_Approver_Admin_Assignment")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("Admin")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Assignment"))
                                    .withNotificationType(NotificationDeliveryMechanism.EMAIL)
                                    .withNotificationLevel(NotificationLevel.CRITICAL)
                                    .withRecipientType(RecipientType.APPROVER)
                                    .withNotificationRecipients(Arrays.asList("approver_admin_member@test.com"))
                                    .withIsDefaultRecipientsEnabled(false),
                                new RoleManagementPolicyExpirationRule()
                                    .withId("Expiration_EndUser_Assignment")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("EndUser")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Assignment"))
                                    .withIsExpirationRequired(true)
                                    .withMaximumDuration("PT7H"),
                                new RoleManagementPolicyEnablementRule()
                                    .withId("Enablement_EndUser_Assignment")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("EndUser")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Assignment"))
                                    .withEnabledRules(
                                        Arrays
                                            .asList(
                                                EnablementRules.JUSTIFICATION,
                                                EnablementRules.MULTI_FACTOR_AUTHENTICATION,
                                                EnablementRules.TICKETING)),
                                new RoleManagementPolicyApprovalRule()
                                    .withId("Approval_EndUser_Assignment")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("EndUser")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Assignment"))
                                    .withSetting(
                                        new ApprovalSettings()
                                            .withIsApprovalRequired(true)
                                            .withIsApprovalRequiredForExtension(false)
                                            .withIsRequestorJustificationRequired(true)
                                            .withApprovalMode(ApprovalMode.SINGLE_STAGE)
                                            .withApprovalStages(
                                                Arrays
                                                    .asList(
                                                        new ApprovalStage()
                                                            .withApprovalStageTimeOutInDays(1)
                                                            .withIsApproverJustificationRequired(true)
                                                            .withEscalationTimeInMinutes(0)
                                                            .withPrimaryApprovers(
                                                                Arrays
                                                                    .asList(
                                                                        new UserSet()
                                                                            .withUserType(UserType.GROUP)
                                                                            .withIsBackup(false)
                                                                            .withId(
                                                                                "2385b0f3-5fa9-43cf-8ca4-b01dc97298cd")
                                                                            .withDescription("amansw_new_group"),
                                                                        new UserSet()
                                                                            .withUserType(UserType.GROUP)
                                                                            .withIsBackup(false)
                                                                            .withId(
                                                                                "2f4913c9-d15b-406a-9946-1d66a28f2690")
                                                                            .withDescription("amansw_group")))
                                                            .withIsEscalationEnabled(false)))),
                                new RoleManagementPolicyAuthenticationContextRule()
                                    .withId("AuthenticationContext_EndUser_Assignment")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("EndUser")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Assignment"))
                                    .withIsEnabled(false)
                                    .withClaimValue(""),
                                new RoleManagementPolicyNotificationRule()
                                    .withId("Notification_Admin_EndUser_Assignment")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("EndUser")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Assignment"))
                                    .withNotificationType(NotificationDeliveryMechanism.EMAIL)
                                    .withNotificationLevel(NotificationLevel.CRITICAL)
                                    .withRecipientType(RecipientType.ADMIN)
                                    .withNotificationRecipients(Arrays.asList("admin_enduser_member@test.com"))
                                    .withIsDefaultRecipientsEnabled(false),
                                new RoleManagementPolicyNotificationRule()
                                    .withId("Notification_Requestor_EndUser_Assignment")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("EndUser")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Assignment"))
                                    .withNotificationType(NotificationDeliveryMechanism.EMAIL)
                                    .withNotificationLevel(NotificationLevel.CRITICAL)
                                    .withRecipientType(RecipientType.REQUESTOR)
                                    .withNotificationRecipients(Arrays.asList("requestor_enduser_member@test.com"))
                                    .withIsDefaultRecipientsEnabled(false),
                                new RoleManagementPolicyNotificationRule()
                                    .withId("Notification_Approver_EndUser_Assignment")
                                    .withTarget(
                                        new RoleManagementPolicyRuleTarget()
                                            .withCaller("EndUser")
                                            .withOperations(Arrays.asList("All"))
                                            .withLevel("Assignment"))
                                    .withNotificationType(NotificationDeliveryMechanism.EMAIL)
                                    .withNotificationLevel(NotificationLevel.CRITICAL)
                                    .withRecipientType(RecipientType.APPROVER)
                                    .withIsDefaultRecipientsEnabled(true))),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
