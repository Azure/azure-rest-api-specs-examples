const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_VirtualMachineConfiguration_ManagedOSDisk.json
 */
async function createPoolVirtualMachineConfigurationOSDisk() {
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
      scaleSettings: { fixedScale: { targetDedicatedNodes: 1, targetLowPriorityNodes: 0 } },
      vmSize: "Standard_d2s_v3",
    },
  );
  console.log(result);
}
