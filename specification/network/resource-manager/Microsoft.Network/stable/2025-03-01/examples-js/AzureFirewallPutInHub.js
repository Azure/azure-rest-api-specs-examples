const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates the specified Azure Firewall.
 *
 * @summary Creates or updates the specified Azure Firewall.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/AzureFirewallPutInHub.json
 */
async function createAzureFirewallInVirtualHub() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const azureFirewallName = "azurefirewall";
  const parameters = {
    firewallPolicy: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/policy1",
    },
    hubIPAddresses: { publicIPs: { addresses: [], count: 1 } },
    location: "West US",
    sku: { name: "AZFW_Hub", tier: "Standard" },
    tags: { key1: "value1" },
    threatIntelMode: "Alert",
    virtualHub: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1",
    },
    zones: [],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.azureFirewalls.beginCreateOrUpdateAndWait(
    resourceGroupName,
    azureFirewallName,
    parameters,
  );
  console.log(result);
}
