const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update a connection monitor.
 *
 * @summary Create or update a connection monitor.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkWatcherConnectionMonitorV2Create.json
 */
async function createConnectionMonitorV2() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const connectionMonitorName = "cm1";
  const options = {
    body: {
      properties: {
        endpoints: [
          {
            name: "vm1",
            resourceId:
              "/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/NwRgIrinaCentralUSEUAP/providers/Microsoft.Compute/virtualMachines/vm1",
          },
          {
            name: "CanaryWorkspaceVamshi",
            filter: {
              type: "Include",
              items: [{ type: "AgentAddress", address: "npmuser" }],
            },
            resourceId:
              "/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/vasamudrRG/providers/Microsoft.OperationalInsights/workspaces/vasamudrWorkspace",
          },
          { name: "bing", address: "bing.com" },
          { name: "google", address: "google.com" },
        ],
        outputs: [],
        testConfigurations: [
          {
            name: "testConfig1",
            tcpConfiguration: { disableTraceRoute: false, port: 80 },
            testFrequencySec: 60,
            protocol: "Tcp",
          },
        ],
        testGroups: [
          {
            name: "test1",
            destinations: ["bing", "google"],
            disable: false,
            sources: ["vm1", "CanaryWorkspaceVamshi"],
            testConfigurations: ["testConfig1"],
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
      connectionMonitorName,
    )
    .put(options);
  const poller = await getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createConnectionMonitorV2().catch(console.error);
