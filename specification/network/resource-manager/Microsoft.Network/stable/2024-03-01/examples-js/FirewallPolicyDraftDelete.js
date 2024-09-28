const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a draft policy.
 *
 * @summary Delete a draft policy.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/FirewallPolicyDraftDelete.json
 */
async function deleteFirewallPolicyDraft() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const firewallPolicyName = "firewallPolicy";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.firewallPolicyDrafts.delete(resourceGroupName, firewallPolicyName);
  console.log(result);
}
