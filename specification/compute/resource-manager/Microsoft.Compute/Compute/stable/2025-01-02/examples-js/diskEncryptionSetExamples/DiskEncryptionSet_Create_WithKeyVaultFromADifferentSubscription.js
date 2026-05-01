const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a disk encryption set
 *
 * @summary creates or updates a disk encryption set
 * x-ms-original-file: 2025-01-02/diskEncryptionSetExamples/DiskEncryptionSet_Create_WithKeyVaultFromADifferentSubscription.json
 */
async function createADiskEncryptionSetWithKeyVaultFromADifferentSubscription() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskEncryptionSets.createOrUpdate(
    "myResourceGroup",
    "myDiskEncryptionSet",
    {
      location: "West US",
      identity: { type: "SystemAssigned" },
      activeKey: { keyUrl: "https://myvaultdifferentsub.vault-int.azure-int.net/keys/{key}" },
      encryptionType: "EncryptionAtRestWithCustomerKey",
    },
  );
  console.log(result);
}
