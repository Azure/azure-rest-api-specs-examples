const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this operation creates or updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 *
 * @summary this operation creates or updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 * x-ms-original-file: 2025-03-01/createPolicyAssignmentWithResourceSelectors.json
 */
async function createOrUpdateAPolicyAssignmentWithResourceSelectors() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyAssignments.create(
    "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2",
    "CostManagement",
    {
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
    },
  );
  console.log(result);
}
