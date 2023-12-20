const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new pool inside the specified account.
 *
 * @summary Creates a new pool inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PoolCreate_AcceleratedNetworking.json
 */
async function createPoolAcceleratedNetworking() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["BATCH_RESOURCE_GROUP"] || "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const parameters = {
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
        "/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123",
    },
    scaleSettings: {
      fixedScale: { targetDedicatedNodes: 1, targetLowPriorityNodes: 0 },
    },
    vmSize: "STANDARD_D1_V2",
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.poolOperations.create(
    resourceGroupName,
    accountName,
    poolName,
    parameters,
  );
  console.log(result);
}
