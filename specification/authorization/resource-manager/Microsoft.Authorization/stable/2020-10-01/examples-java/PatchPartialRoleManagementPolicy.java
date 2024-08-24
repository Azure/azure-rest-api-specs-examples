
import com.azure.resourcemanager.authorization.fluent.models.RoleManagementPolicyInner;
import com.azure.resourcemanager.authorization.models.NotificationDeliveryMechanism;
import com.azure.resourcemanager.authorization.models.NotificationLevel;
import com.azure.resourcemanager.authorization.models.RecipientType;
import com.azure.resourcemanager.authorization.models.RoleManagementPolicyExpirationRule;
import com.azure.resourcemanager.authorization.models.RoleManagementPolicyNotificationRule;
import com.azure.resourcemanager.authorization.models.RoleManagementPolicyRuleTarget;
import java.util.Arrays;

/**
 * Samples for RoleManagementPolicies Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/
     * PatchPartialRoleManagementPolicy.json
     */
    /**
     * Sample code: PatchPartialRoleManagementPolicy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchPartialRoleManagementPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleManagementPolicies()
            .updateWithResponse("providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368",
                "570c3619-7688-4b34-b290-2b8bb3ccab2a",
                new RoleManagementPolicyInner().withRules(Arrays.asList(
                    new RoleManagementPolicyExpirationRule().withId("Expiration_Admin_Eligibility")
                        .withTarget(new RoleManagementPolicyRuleTarget().withCaller("Admin")
                            .withOperations(Arrays.asList("All")).withLevel("Eligibility"))
                        .withIsExpirationRequired(false).withMaximumDuration("P180D"),
                    new RoleManagementPolicyNotificationRule().withId("Notification_Admin_Admin_Eligibility")
                        .withTarget(new RoleManagementPolicyRuleTarget().withCaller("Admin")
                            .withOperations(Arrays.asList("All")).withLevel("Eligibility"))
                        .withNotificationType(NotificationDeliveryMechanism.EMAIL)
                        .withNotificationLevel(NotificationLevel.CRITICAL).withRecipientType(RecipientType.ADMIN)
                        .withNotificationRecipients(Arrays.asList("admin_admin_eligible@test.com"))
                        .withIsDefaultRecipientsEnabled(false))),
                com.azure.core.util.Context.NONE);
    }
}
