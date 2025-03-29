const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to  This operation creates or updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 *
 * @summary  This operation creates or updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/createPolicyAssignmentWithResourceSelectors.json
 */
async function createOrUpdateAPolicyAssignmentWithResourceSelectors() {
  const scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const policyAssignmentName = "CostManagement";
  const parameters = {
    description: "Limit the resource location and resource SKU",
    displayName: "Limit the resource location and resource SKU",
    metadata: { assignedBy: "Special Someone" },
    policyDefinitionId:
      "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement",
    resourceSelectors: [
      {
        name: "SDPRegions",
        selectors: [{ in: ["eastus2euap", "centraluseuap"], kind: "resourceLocation" }],
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyAssignments.create(scope, policyAssignmentName, parameters);
  console.log(result);
}
