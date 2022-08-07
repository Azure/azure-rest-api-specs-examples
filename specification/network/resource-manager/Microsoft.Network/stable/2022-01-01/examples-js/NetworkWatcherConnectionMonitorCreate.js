const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a connection monitor.
 *
 * @summary Create or update a connection monitor.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkWatcherConnectionMonitorCreate.json
 */
async function createConnectionMonitorV1() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
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
    parameters
  );
  console.log(result);
}

createConnectionMonitorV1().catch(console.error);
