const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a role management policy
 *
 * @summary update a role management policy
 * x-ms-original-file: 2024-09-01-preview/PatchPartialRoleManagementPolicy.json
 */
async function patchPartialRoleManagementPolicy() {
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.roleManagementPolicies.update(
    "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368",
    "570c3619-7688-4b34-b290-2b8bb3ccab2a",
    {
      rules: [
        {
          id: "Expiration_Admin_Eligibility",
          isExpirationRequired: false,
          maximumDuration: "P180D",
          ruleType: "RoleManagementPolicyExpirationRule",
          target: { caller: "Admin", level: "Eligibility", operations: ["All"] },
        },
        {
          id: "Notification_Admin_Admin_Eligibility",
          isDefaultRecipientsEnabled: false,
          notificationLevel: "Critical",
          notificationRecipients: ["admin_admin_eligible@test.com"],
          notificationType: "Email",
          recipientType: "Admin",
          ruleType: "RoleManagementPolicyNotificationRule",
          target: { caller: "Admin", level: "Eligibility", operations: ["All"] },
        },
      ],
    },
  );
  console.log(result);
}
