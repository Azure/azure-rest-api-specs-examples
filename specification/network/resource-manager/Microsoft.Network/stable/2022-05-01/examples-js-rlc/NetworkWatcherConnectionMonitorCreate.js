const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Create or update a connection monitor.
 *
 * @summary Create or update a connection monitor.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkWatcherConnectionMonitorCreate.json
 */
async function createConnectionMonitorV1() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const connectionMonitorName = "cm1";
  const options = {
    body: {
      location: "eastus",
      properties: {
        endpoints: [
          {
            name: "source",
            resourceId:
              "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1",
          },
          { name: "destination", address: "bing.com" },
        ],
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
      },
    },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}",
      subscriptionId,
      resourceGroupName,
      networkWatcherName,
      connectionMonitorName
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createConnectionMonitorV1().catch(console.error);
