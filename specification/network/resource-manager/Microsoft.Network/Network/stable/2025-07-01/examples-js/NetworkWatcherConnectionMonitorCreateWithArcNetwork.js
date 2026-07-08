const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a connection monitor.
 *
 * @summary create or update a connection monitor.
 * x-ms-original-file: 2025-07-01/NetworkWatcherConnectionMonitorCreateWithArcNetwork.json
 */
async function createConnectionMonitorWithArcNetwork() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.connectionMonitors.createOrUpdate("rg1", "nw1", "cm1", {
    endpoints: [
      {
        name: "vm1",
        type: "AzureVM",
        resourceId:
          "/subscriptions/9cece3e3-0f7d-47ca-af0e-9772773f90b7/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/TESTVM",
      },
      { name: "bing", type: "ExternalAddress", address: "bing.com" },
      { name: "google", type: "ExternalAddress", address: "google.com" },
      {
        name: "ArcBasedNetwork",
        type: "AzureArcNetwork",
        locationDetails: { region: "eastus" },
        scope: { include: [{ address: "172.21.128.0/20" }] },
        subscriptionId: "9cece3e3-0f7d-47ca-af0e-9772773f90b7",
      },
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
        sources: ["vm1", "ArcBasedNetwork"],
        testConfigurations: ["testConfig1"],
      },
    ],
  });
  console.log(result);
}
