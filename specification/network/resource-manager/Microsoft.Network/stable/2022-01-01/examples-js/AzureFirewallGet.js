const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified Azure Firewall.
 *
 * @summary Gets the specified Azure Firewall.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/AzureFirewallGet.json
 */
async function getAzureFirewall() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const azureFirewallName = "azurefirewall";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.azureFirewalls.get(resourceGroupName, azureFirewallName);
  console.log(result);
}

getAzureFirewall().catch(console.error);
