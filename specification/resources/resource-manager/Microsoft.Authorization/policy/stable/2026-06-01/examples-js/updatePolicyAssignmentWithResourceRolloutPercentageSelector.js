const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this operation updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 *
 * @summary this operation updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 * x-ms-original-file: 2026-06-01/updatePolicyAssignmentWithResourceRolloutPercentageSelector.json
 */
async function updateAPolicyAssignmentWithAResourceRolloutPercentageSelector() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyAssignments.update(
    "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2",
    "CostManagement",
    {
      resourceSelectors: [
        { name: "SDPRollout", selectors: [{ kind: "resourceRolloutPercentage", progress: 80 }] },
      ],
    },
  );
  console.log(result);
}
