const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates (patches) a disk encryption set.
 *
 * @summary Updates (patches) a disk encryption set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/UpdateADiskEncryptionSetWithRotationToLatestKeyVersionEnabled.json
 */
async function updateADiskEncryptionSetWithRotationToLatestKeyVersionEnabledSetToTrueSucceeded() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskEncryptionSetName = "myDiskEncryptionSet";
  const diskEncryptionSet = {
    activeKey: {
      keyUrl: "https://myvaultdifferentsub.vault-int.azure-int.net/keys/keyName/keyVersion1",
    },
    encryptionType: "EncryptionAtRestWithCustomerKey",
    identity: { type: "SystemAssigned" },
    rotationToLatestKeyVersionEnabled: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskEncryptionSets.beginUpdateAndWait(
    resourceGroupName,
    diskEncryptionSetName,
    diskEncryptionSet
  );
  console.log(result);
}

updateADiskEncryptionSetWithRotationToLatestKeyVersionEnabledSetToTrueSucceeded().catch(
  console.error
);
