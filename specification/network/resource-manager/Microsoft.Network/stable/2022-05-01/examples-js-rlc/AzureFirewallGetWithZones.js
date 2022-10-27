const createNetworkManagementClient = require("@azure-rest/arm-network").default;
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Gets the specified Azure Firewall.
 *
 * @summary Gets the specified Azure Firewall.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/AzureFirewallGetWithZones.json
 */
async function getAzureFirewallWithZones() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const azureFirewallName = "azurefirewall";
  const options = {
    queryParameters: { "api-version": "2022-05-01" },
  };
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}",
      subscriptionId,
      resourceGroupName,
      azureFirewallName
    )
    .get(options);
  console.log(result);
}

getAzureFirewallWithZones().catch(console.error);
