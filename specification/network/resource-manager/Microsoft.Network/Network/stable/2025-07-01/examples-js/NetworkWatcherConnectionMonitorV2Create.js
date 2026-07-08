const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a connection monitor.
 *
 * @summary create or update a connection monitor.
 * x-ms-original-file: 2025-07-01/NetworkWatcherConnectionMonitorV2Create.json
 */
async function createConnectionMonitorV2() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.connectionMonitors.createOrUpdate("rg1", "nw1", "cm1", {
    endpoints: [
      {
        name: "vm1",
        resourceId:
          "/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/NwRgIrinaCentralUSEUAP/providers/Microsoft.Compute/virtualMachines/vm1",
      },
      {
        name: "CanaryWorkspaceVamshi",
        filter: { type: "Include", items: [{ type: "AgentAddress", address: "npmuser" }] },
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
  });
  console.log(result);
}
