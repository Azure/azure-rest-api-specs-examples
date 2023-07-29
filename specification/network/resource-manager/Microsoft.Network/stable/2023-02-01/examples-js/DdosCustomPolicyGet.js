const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified DDoS custom policy.
 *
 * @summary Gets information about the specified DDoS custom policy.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/DdosCustomPolicyGet.json
 */
async function getDDoSCustomPolicy() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const ddosCustomPolicyName = "test-ddos-custom-policy";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ddosCustomPolicies.get(resourceGroupName, ddosCustomPolicyName);
  console.log(result);
}
