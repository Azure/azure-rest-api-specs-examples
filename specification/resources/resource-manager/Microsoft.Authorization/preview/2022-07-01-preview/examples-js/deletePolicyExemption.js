const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation deletes a policy exemption, given its name and the scope it was created in. The scope of a policy exemption is the part of its ID preceding '/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}'.
 *
 * @summary This operation deletes a policy exemption, given its name and the scope it was created in. The scope of a policy exemption is the part of its ID preceding '/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}'.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/deletePolicyExemption.json
 */
async function deleteAPolicyExemption() {
  const scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster";
  const policyExemptionName = "DemoExpensiveVM";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyExemptions.delete(scope, policyExemptionName);
  console.log(result);
}
