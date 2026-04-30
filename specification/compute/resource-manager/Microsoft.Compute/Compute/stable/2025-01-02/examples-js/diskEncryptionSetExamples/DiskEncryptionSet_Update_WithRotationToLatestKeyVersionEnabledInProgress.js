const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates (patches) a disk encryption set.
 *
 * @summary updates (patches) a disk encryption set.
 * x-ms-original-file: 2025-01-02/diskEncryptionSetExamples/DiskEncryptionSet_Update_WithRotationToLatestKeyVersionEnabledInProgress.json
 */
async function updateADiskEncryptionSetWithRotationToLatestKeyVersionEnabledSetToTrueUpdating() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskEncryptionSets.update("myResourceGroup", "myDiskEncryptionSet", {
    identity: { type: "SystemAssigned" },
    activeKey: {
      keyUrl: "https://myvaultdifferentsub.vault-int.azure-int.net/keys/keyName/keyVersion1",
    },
    encryptionType: "EncryptionAtRestWithCustomerKey",
    rotationToLatestKeyVersionEnabled: true,
  });
  console.log(result);
}
