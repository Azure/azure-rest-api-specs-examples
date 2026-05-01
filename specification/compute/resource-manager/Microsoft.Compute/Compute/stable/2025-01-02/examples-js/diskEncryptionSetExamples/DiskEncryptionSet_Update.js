const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates (patches) a disk encryption set.
 *
 * @summary updates (patches) a disk encryption set.
 * x-ms-original-file: 2025-01-02/diskEncryptionSetExamples/DiskEncryptionSet_Update.json
 */
async function updateADiskEncryptionSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskEncryptionSets.update("myResourceGroup", "myDiskEncryptionSet", {
    activeKey: {
      sourceVault: {
        id: "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault",
      },
      keyUrl: "https://myvmvault.vault-int.azure-int.net/keys/keyName/keyVersion",
    },
    encryptionType: "EncryptionAtRestWithCustomerKey",
    tags: { department: "Development", project: "Encryption" },
  });
  console.log(result);
}
