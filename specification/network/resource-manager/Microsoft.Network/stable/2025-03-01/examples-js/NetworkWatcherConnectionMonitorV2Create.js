const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update a connection monitor.
 *
 * @summary Create or update a connection monitor.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkWatcherConnectionMonitorV2Create.json
 */
async function createConnectionMonitorV2() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkWatcherName = "nw1";
  const connectionMonitorName = "cm1";
  const parameters = {
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
