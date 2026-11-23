const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified Network Virtual Appliance.
 *
 * @summary creates or updates the specified Network Virtual Appliance.
 * x-ms-original-file: 2026-05-01/NetworkVirtualApplianceVhubSourceIPProtocolPut.json
 */
async function createOrUpdateAManagedNVAInAVirtualHubWithSourceIPProtocolLoadDistribution() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkVirtualAppliances.createOrUpdate("rg1", "nva-vhub", {
    location: "West US",
    nvaSku: { bundledScaleUnit: "1", marketPlaceVersion: "12.1", vendor: "Cisco SDWAN" },
    virtualHub: {
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1",
    },
    addressFamily: ["IPv4"],
    virtualApplianceAsn: 10000,
    loadBalancerConfiguration: { loadDistribution: "SourceIPProtocol", idleTimeoutInMinutes: 15 },
  });
  console.log(result);
}
