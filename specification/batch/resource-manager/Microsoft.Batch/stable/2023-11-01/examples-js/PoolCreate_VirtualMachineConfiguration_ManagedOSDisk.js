const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new pool inside the specified account.
 *
 * @summary Creates a new pool inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PoolCreate_VirtualMachineConfiguration_ManagedOSDisk.json
 */
async function createPoolVirtualMachineConfigurationOSDisk() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["BATCH_RESOURCE_GROUP"] || "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const parameters = {
    deploymentConfiguration: {
      virtualMachineConfiguration: {
        imageReference: {
          offer: "windowsserver",
          publisher: "microsoftwindowsserver",
          sku: "2022-datacenter-smalldisk",
        },
        nodeAgentSkuId: "batch.node.windows amd64",
        osDisk: {
          caching: "ReadWrite",
          diskSizeGB: 100,
          managedDisk: { storageAccountType: "StandardSSD_LRS" },
          writeAcceleratorEnabled: false,
        },
      },
    },
    scaleSettings: {
      fixedScale: { targetDedicatedNodes: 1, targetLowPriorityNodes: 0 },
    },
    vmSize: "Standard_d2s_v3",
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
