const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this operation creates or updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 *
 * @summary this operation creates or updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 * x-ms-original-file: 2026-01-01-preview/createPolicyAssignmentWithSelfserveExemptionSettings.json
 */
async function createOrUpdateAPolicyAssignmentWithSelfServeExemptionSettings() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyAssignments.create(
    "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2",
    "CostManagement",
    {
      location: "eastus",
      description: "Limit the resource location and resource SKU",
      definitionVersion: "1.*.*",
      displayName: "Limit the resource location and resource SKU",
      enforcementMode: "Default",
      metadata: { assignedBy: "Foo Bar" },
      policyDefinitionId:
        "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement",
      selfServeExemptionSettings: { enabled: true, policyDefinitionReferenceIds: ["Limit_Skus"] },
    },
  );
  console.log(result);
}
