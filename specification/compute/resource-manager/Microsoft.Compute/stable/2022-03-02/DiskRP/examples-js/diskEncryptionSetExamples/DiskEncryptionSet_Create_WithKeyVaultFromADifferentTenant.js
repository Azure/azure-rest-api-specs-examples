const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a disk encryption set
 *
 * @summary Creates or updates a disk encryption set
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-02/DiskRP/examples/diskEncryptionSetExamples/DiskEncryptionSet_Create_WithKeyVaultFromADifferentTenant.json
 */
async function createADiskEncryptionSetWithKeyVaultFromADifferentTenant() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskEncryptionSetName = "myDiskEncryptionSet";
  const diskEncryptionSet = {
    activeKey: {
      keyUrl: "https://myvaultdifferenttenant.vault-int.azure-int.net/keys/{key}",
    },
    encryptionType: "EncryptionAtRestWithCustomerKey",
    federatedClientId: "00000000-0000-0000-0000-000000000000",
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/MicrosoftManagedIdentity/userAssignedIdentities/{identityName}":
          {},
      },
    },
    location: "West US",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskEncryptionSets.beginCreateOrUpdateAndWait(
    resourceGroupName,
    diskEncryptionSetName,
    diskEncryptionSet
  );
  console.log(result);
}

createADiskEncryptionSetWithKeyVaultFromADifferentTenant().catch(console.error);
