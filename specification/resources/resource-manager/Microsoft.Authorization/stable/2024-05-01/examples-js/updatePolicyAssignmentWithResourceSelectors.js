const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to  This operation updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 *
 * @summary  This operation updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2024-05-01/examples/updatePolicyAssignmentWithResourceSelectors.json
 */
async function updateAPolicyAssignmentWithResourceSelectors() {
  const scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const policyAssignmentName = "CostManagement";
  const parameters = {
    resourceSelectors: [
      {
        name: "SDPRegions",
        selectors: [{ in: ["eastus2euap", "centraluseuap"], kind: "resourceLocation" }],
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyAssignments.update(scope, policyAssignmentName, parameters);
  console.log(result);
}
