const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves a list of all IP prefixes that azure firewall has learned to not SNAT.
 *
 * @summary Retrieves a list of all IP prefixes that azure firewall has learned to not SNAT.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/AzureFirewallListLearnedIPPrefixes.json
 */
async function azureFirewallListLearnedPrefixes() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const azureFirewallName = "azureFirewall1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.azureFirewalls.beginListLearnedPrefixesAndWait(
    resourceGroupName,
    azureFirewallName
  );
  console.log(result);
}

azureFirewallListLearnedPrefixes().catch(console.error);
