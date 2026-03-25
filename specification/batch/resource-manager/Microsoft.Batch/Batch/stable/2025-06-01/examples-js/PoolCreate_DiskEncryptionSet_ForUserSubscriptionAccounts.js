const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_DiskEncryptionSet_ForUserSubscriptionAccounts.json
 */
async function createPoolDiskEncryptionSetForUserSubscriptionAccounts() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.pool.create(
    "default-azurebatch-japaneast",
    "sampleacct",
    "testpool",
    {
      vmSize: "Standard_D4ds_v5",
      taskSchedulingPolicy: { nodeFillType: "Pack" },
      deploymentConfiguration: {
        virtualMachineConfiguration: {
          imageReference: {
            sku: "2022-Datacenter",
            publisher: "MicrosoftWindowsServer",
            version: "latest",
            offer: "WindowsServer",
          },
          nodeAgentSkuId: "batch.node.windows amd64",
          securityProfile: { encryptionAtHost: false },
          osDisk: {
            managedDisk: {
              storageAccountType: "Standard_LRS",
              diskEncryptionSet: {
                id: "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/diskEncryptionSets/DiskEncryptionSetId",
              },
            },
          },
        },
      },
      scaleSettings: { fixedScale: { targetDedicatedNodes: 1, resizeTimeout: "PT15M" } },
    },
  );
  console.log(result);
}
