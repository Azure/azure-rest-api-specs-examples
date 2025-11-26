const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates the specified Azure Firewall.
 *
 * @summary Creates or updates the specified Azure Firewall.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/AzureFirewallPutWithZones.json
 */
async function createAzureFirewallWithZones() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const azureFirewallName = "azurefirewall";
  const parameters = {
    applicationRuleCollections: [
      {
        name: "apprulecoll",
        action: { type: "Deny" },
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/applicationRuleCollections/apprulecoll",
        priority: 110,
        rules: [
          {
            name: "rule1",
            description: "Deny inbound rule",
            protocols: [{ port: 443, protocolType: "Https" }],
            sourceAddresses: ["216.58.216.164", "10.0.0.0/24"],
            targetFqdns: ["www.test.com"],
          },
        ],
      },
    ],
    ipConfigurations: [
      {
        name: "azureFirewallIpConfiguration",
        publicIPAddress: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName",
        },
        subnet: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/AzureFirewallSubnet",
        },
      },
    ],
    location: "West US 2",
    natRuleCollections: [
      {
        name: "natrulecoll",
        action: { type: "Dnat" },
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/natRuleCollections/natrulecoll",
        priority: 112,
        rules: [
          {
            name: "DNAT-HTTPS-traffic",
            description: "D-NAT all outbound web traffic for inspection",
            destinationAddresses: ["1.2.3.4"],
            destinationPorts: ["443"],
            protocols: ["TCP"],
            sourceAddresses: ["*"],
            translatedAddress: "1.2.3.5",
            translatedPort: "8443",
          },
          {
            name: "DNAT-HTTP-traffic-With-FQDN",
            description: "D-NAT all inbound web traffic for inspection",
            destinationAddresses: ["1.2.3.4"],
            destinationPorts: ["80"],
            protocols: ["TCP"],
            sourceAddresses: ["*"],
            translatedFqdn: "internalhttpserver",
            translatedPort: "880",
          },
        ],
      },
    ],
    networkRuleCollections: [
      {
        name: "netrulecoll",
        action: { type: "Deny" },
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/networkRuleCollections/netrulecoll",
        priority: 112,
        rules: [
          {
            name: "L4-traffic",
            description: "Block traffic based on source IPs and ports",
            destinationAddresses: ["*"],
            destinationPorts: ["443-444", "8443"],
            protocols: ["TCP"],
            sourceAddresses: ["192.168.1.1-192.168.1.12", "10.1.4.12-10.1.4.255"],
          },
          {
            name: "L4-traffic-with-FQDN",
            description: "Block traffic based on source IPs and ports to amazon",
            destinationFqdns: ["www.amazon.com"],
            destinationPorts: ["443-444", "8443"],
            protocols: ["TCP"],
            sourceAddresses: ["10.2.4.12-10.2.4.255"],
          },
        ],
      },
    ],
    sku: { name: "AZFW_VNet", tier: "Standard" },
    tags: { key1: "value1" },
    threatIntelMode: "Alert",
    zones: ["1", "2", "3"],
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
