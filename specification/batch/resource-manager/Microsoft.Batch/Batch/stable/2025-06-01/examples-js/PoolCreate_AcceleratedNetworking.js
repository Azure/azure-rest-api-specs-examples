const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_AcceleratedNetworking.json
 */
async function createPoolAcceleratedNetworking() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.pool.create(
    "default-azurebatch-japaneast",
    "sampleacct",
    "testpool",
    {
      deploymentConfiguration: {
        virtualMachineConfiguration: {
          imageReference: {
            offer: "WindowsServer",
            publisher: "MicrosoftWindowsServer",
            sku: "2016-datacenter-smalldisk",
            version: "latest",
          },
          nodeAgentSkuId: "batch.node.windows amd64",
        },
      },
      networkConfiguration: {
        enableAcceleratedNetworking: true,
        subnetId:
          "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123",
      },
      scaleSettings: { fixedScale: { targetDedicatedNodes: 1, targetLowPriorityNodes: 0 } },
      vmSize: "STANDARD_D1_V2",
    },
  );
  console.log(result);
}
