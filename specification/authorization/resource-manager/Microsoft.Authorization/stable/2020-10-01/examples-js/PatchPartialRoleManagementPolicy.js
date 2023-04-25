const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a role management policy
 *
 * @summary Update a role management policy
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/PatchPartialRoleManagementPolicy.json
 */
async function patchPartialRoleManagementPolicy() {
  const subscriptionId =
    process.env["AUTHORIZATION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope =
    "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368";
  const roleManagementPolicyName = "570c3619-7688-4b34-b290-2b8bb3ccab2a";
  const parameters = {
    rules: [
      {
        id: "Expiration_Admin_Eligibility",
        isExpirationRequired: false,
        maximumDuration: "P180D",
        ruleType: "RoleManagementPolicyExpirationRule",
        target: {
          caller: "Admin",
          enforcedSettings: [],
          inheritableSettings: [],
          level: "Eligibility",
          targetObjects: [],
          operations: ["All"],
        },
      },
      {
        id: "Notification_Admin_Admin_Eligibility",
        isDefaultRecipientsEnabled: false,
        notificationLevel: "Critical",
        notificationRecipients: ["admin_admin_eligible@test.com"],
        notificationType: "Email",
        recipientType: "Admin",
        ruleType: "RoleManagementPolicyNotificationRule",
        target: {
          caller: "Admin",
          enforcedSettings: [],
          inheritableSettings: [],
          level: "Eligibility",
          targetObjects: [],
          operations: ["All"],
        },
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential, subscriptionId);
  const result = await client.roleManagementPolicies.update(
    scope,
    roleManagementPolicyName,
    parameters
  );
  console.log(result);
}
