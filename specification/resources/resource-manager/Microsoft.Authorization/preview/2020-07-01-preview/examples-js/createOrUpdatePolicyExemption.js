const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to  This operation creates or updates a policy exemption with the given scope and name. Policy exemptions apply to all resources contained within their scope. For example, when you create a policy exemption at resource group scope for a policy assignment at the same or above level, the exemption exempts to all applicable resources in the resource group.
 *
 * @summary  This operation creates or updates a policy exemption with the given scope and name. Policy exemptions apply to all resources contained within their scope. For example, when you create a policy exemption at resource group scope for a policy assignment at the same or above level, the exemption exempts to all applicable resources in the resource group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/createOrUpdatePolicyExemption.json
 */
async function createOrUpdateAPolicyExemption() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster";
  const policyExemptionName = "DemoExpensiveVM";
  const parameters = {
    description: "Exempt demo cluster from limit sku",
    displayName: "Exempt demo cluster",
    exemptionCategory: "Waiver",
    metadata: { reason: "Temporary exemption for a expensive VM demo" },
    policyAssignmentId:
      "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement",
    policyDefinitionReferenceIds: ["Limit_Skus"],
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policyExemptions.createOrUpdate(
    scope,
    policyExemptionName,
    parameters,
  );
  console.log(result);
}
