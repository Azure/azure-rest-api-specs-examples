const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified Azure Firewall.
 *
 * @summary gets the specified Azure Firewall.
 * x-ms-original-file: 2025-07-01/AzureFirewallGetWithMgmtSubnet.json
 */
async function getAzureFirewallWithManagementSubnet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.azureFirewalls.get("rg1", "azurefirewall");
  console.log(result);
}
