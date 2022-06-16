const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified DDoS custom policy.
 *
 * @summary Gets information about the specified DDoS custom policy.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/DdosCustomPolicyGet.json
 */
async function getDDoSCustomPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const ddosCustomPolicyName = "test-ddos-custom-policy";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ddosCustomPolicies.get(resourceGroupName, ddosCustomPolicyName);
  console.log(result);
}

getDDoSCustomPolicy().catch(console.error);
