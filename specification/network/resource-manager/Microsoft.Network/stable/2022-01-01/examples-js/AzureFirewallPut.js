const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the specified Azure Firewall.
 *
 * @summary Creates or updates the specified Azure Firewall.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/AzureFirewallPut.json
 */
async function createAzureFirewall() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const azureFirewallName = "azurefirewall";
  const parameters = {
    applicationRuleCollections: [
      {
        name: "apprulecoll",
        action: { type: "Deny" },
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
    location: "West US",
    natRuleCollections: [
      {
        name: "natrulecoll",
        action: { type: "Dnat" },
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
    zones: [],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.azureFirewalls.beginCreateOrUpdateAndWait(
    resourceGroupName,
    azureFirewallName,
    parameters
  );
  console.log(result);
}

createAzureFirewall().catch(console.error);
