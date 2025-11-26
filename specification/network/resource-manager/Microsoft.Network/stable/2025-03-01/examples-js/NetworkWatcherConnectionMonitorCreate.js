const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update a connection monitor.
 *
 * @summary Create or update a connection monitor.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkWatcherConnectionMonitorCreate.json
 */
async function createConnectionMonitorV1() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkWatcherName = "nw1";
  const connectionMonitorName = "cm1";
  const parameters = {
    endpoints: [
      {
        name: "source",
        resourceId:
          "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1",
      },
      { name: "destination", address: "bing.com" },
    ],
    location: "eastus",
    testConfigurations: [
      {
        name: "tcp",
        tcpConfiguration: { port: 80 },
        testFrequencySec: 60,
        protocol: "Tcp",
      },
    ],
    testGroups: [
      {
        name: "tg",
        destinations: ["destination"],
        sources: ["source"],
        testConfigurations: ["tcp"],
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.connectionMonitors.beginCreateOrUpdateAndWait(
    resourceGroupName,
    networkWatcherName,
    connectionMonitorName,
    parameters,
  );
  console.log(result);
}
