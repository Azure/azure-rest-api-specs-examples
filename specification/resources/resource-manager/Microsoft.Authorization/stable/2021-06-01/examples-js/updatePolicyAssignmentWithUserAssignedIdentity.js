const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to  This operation updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 *
 * @summary  This operation updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/updatePolicyAssignmentWithUserAssignedIdentity.json
 */
async function updateAPolicyAssignmentWithAUserAssignedIdentity() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const policyAssignmentName = "EnforceNaming";
  const parameters = {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/ae640e6bBa3e42569d622993eecfa6f2/resourceGroups/testResourceGroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/testIdentity":
          {},
      },
    },
    location: "eastus",
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policyAssignments.update(scope, policyAssignmentName, parameters);
  console.log(result);
}
