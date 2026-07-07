const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a connection monitor.
 *
 * @summary create or update a connection monitor.
 * x-ms-original-file: 2025-07-01/NetworkWatcherConnectionMonitorCreate.json
 */
async function createConnectionMonitorV1() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.connectionMonitors.createOrUpdate("rg1", "nw1", "cm1", {
    location: "eastus",
    endpoints: [
      {
        name: "source",
        resourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1",
      },
      { name: "destination", address: "bing.com" },
    ],
    testConfigurations: [
      { name: "tcp", tcpConfiguration: { port: 80 }, testFrequencySec: 60, protocol: "Tcp" },
    ],
    testGroups: [
      {
        name: "tg",
        destinations: ["destination"],
        sources: ["source"],
        testConfigurations: ["tcp"],
      },
    ],
  });
  console.log(result);
}
