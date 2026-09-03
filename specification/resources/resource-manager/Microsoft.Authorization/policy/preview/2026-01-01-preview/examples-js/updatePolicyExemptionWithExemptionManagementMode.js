const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this operation updates a policy exemption with the given scope and name.
 *
 * @summary this operation updates a policy exemption with the given scope and name.
 * x-ms-original-file: 2026-01-01-preview/updatePolicyExemptionWithExemptionManagementMode.json
 */
async function updateAPolicyExemptionWithExemptionManagementMode() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyExemptions.update(
    "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster",
    "DemoExpensiveVM",
    { exemptionManagementMode: "Admin" },
  );
  console.log(result);
}
