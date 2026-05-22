const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a role management policy
 *
 * @summary update a role management policy
 * x-ms-original-file: 2024-09-01-preview/PatchRoleManagementPolicyToEnablePIMOnlyMode.json
 */
async function patchRoleManagementPolicyToEnablePIMOnlyMode() {
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.roleManagementPolicies.update(
    "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368",
    "570c3619-7688-4b34-b290-2b8bb3ccab2a",
    {
      rules: [
        {
          id: "PIMOnlyMode_Admin_Assignment",
          pimOnlyModeSettings: {
            excludedAssignmentTypes: ["ServicePrincipalsAsTarget"],
            excludes: [
              { type: "User", id: "ec42a424-a0c0-4418-8788-d19bdeb03704" },
              { type: "Group", id: "00029dfb-0218-4e7a-9a85-c15dc0c880bc" },
              { type: "ServicePrincipal", id: "0000103d-1fc2-4ac8-81de-71517765655c" },
            ],
            mode: "Enabled",
          },
          ruleType: "RoleManagementPolicyPimOnlyModeRule",
          target: {
            caller: "Admin",
            enforcedSettings: ["all"],
            inheritableSettings: ["all"],
            level: "Assignment",
            targetObjects: [],
            operations: ["all"],
          },
        },
      ],
    },
  );
  console.log(result);
}
