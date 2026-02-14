const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update a connection monitor.
 *
 * @summary Create or update a connection monitor.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/Network/stable/2025-05-01/examples/NetworkWatcherConnectionMonitorCreateWithArcNetwork.json
 */
async function createConnectionMonitorWithArcNetwork() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkWatcherName = "nw1";
  const connectionMonitorName = "cm1";
  const parameters = {
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
