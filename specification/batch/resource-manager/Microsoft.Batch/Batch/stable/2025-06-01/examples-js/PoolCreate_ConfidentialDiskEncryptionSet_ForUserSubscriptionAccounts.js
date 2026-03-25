const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_ConfidentialDiskEncryptionSet_ForUserSubscriptionAccounts.json
 */
async function createPoolConfidentialDiskEncryptionSetForUserSubscriptionAccounts() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.pool.create(
    "default-azurebatch-japaneast",
    "sampleacct",
    "testpool",
    {
      vmSize: "Standard_DC2as_v5",
      taskSchedulingPolicy: { nodeFillType: "Pack" },
      deploymentConfiguration: {
        virtualMachineConfiguration: {
          imageReference: {
            publisher: "MicrosoftWindowsServer",
            offer: "WindowsServer",
            sku: "2019-datacenter-core-g2",
            version: "latest",
          },
          nodeAgentSkuId: "batch.node.windows amd64",
          securityProfile: {
            securityType: "confidentialVM",
            encryptionAtHost: false,
            uefiSettings: { vTpmEnabled: true, secureBootEnabled: true },
          },
          osDisk: {
            managedDisk: {
              storageAccountType: "Standard_LRS",
              diskEncryptionSet: {
                id: "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/diskEncryptionSets/DiskEncryptionSetId",
              },
              securityProfile: {
                securityEncryptionType: "DiskWithVMGuestState",
                diskEncryptionSet: {
                  id: "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/diskEncryptionSets/ConfidentialDiskEncryptionSetId",
                },
              },
            },
          },
          dataDisks: [
            {
              lun: 0,
              diskSizeGB: 1024,
              managedDisk: {
                storageAccountType: "Standard_LRS",
                diskEncryptionSet: {
                  id: "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/diskEncryptionSets/DiskEncryptionSetId",
                },
              },
            },
          ],
        },
      },
      scaleSettings: { fixedScale: { targetDedicatedNodes: 1, resizeTimeout: "PT15M" } },
    },
  );
  console.log(result);
}
