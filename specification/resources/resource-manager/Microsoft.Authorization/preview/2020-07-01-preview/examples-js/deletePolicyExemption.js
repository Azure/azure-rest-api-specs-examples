const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation deletes a policy exemption, given its name and the scope it was created in. The scope of a policy exemption is the part of its ID preceding '/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}'.
 *
 * @summary This operation deletes a policy exemption, given its name and the scope it was created in. The scope of a policy exemption is the part of its ID preceding '/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}'.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/deletePolicyExemption.json
 */
async function deleteAPolicyExemption() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster";
  const policyExemptionName = "DemoExpensiveVM";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policyExemptions.delete(scope, policyExemptionName);
  console.log(result);
}

deleteAPolicyExemption().catch(console.error);
