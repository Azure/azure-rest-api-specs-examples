const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes Policy
 *
 * @summary Deletes Policy
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyDelete.json
 */
async function deleteProtectionPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const policyName = "Policy1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.policies.delete(resourceGroupName, policyName);
  console.log(result);
}

deleteProtectionPolicy().catch(console.error);
