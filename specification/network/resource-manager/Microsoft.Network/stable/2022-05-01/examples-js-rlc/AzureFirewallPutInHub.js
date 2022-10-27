const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Creates or updates the specified Azure Firewall.
 *
 * @summary Creates or updates the specified Azure Firewall.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/AzureFirewallPutInHub.json
 */
async function createAzureFirewallInVirtualHub() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const azureFirewallName = "azurefirewall";
  const options = {
    body: {
      location: "West US",
      properties: {
        firewallPolicy: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/policy1",
        },
        hubIPAddresses: { publicIPs: { addresses: [], count: 1 } },
        sku: { name: "AZFW_Hub", tier: "Standard" },
        threatIntelMode: "Alert",
        virtualHub: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1",
        },
      },
      tags: { key1: "value1" },
      zones: [],
    },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}",
      subscriptionId,
      resourceGroupName,
      azureFirewallName
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createAzureFirewallInVirtualHub().catch(console.error);
