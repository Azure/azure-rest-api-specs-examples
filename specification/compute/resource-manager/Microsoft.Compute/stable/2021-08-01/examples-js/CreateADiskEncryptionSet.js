const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function createADiskEncryptionSet() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskEncryptionSetName = "myDiskEncryptionSet";
  const diskEncryptionSet = {
    activeKey: {
      keyUrl: "https://myvmvault.vault-int.azure-int.net/keys/{key}",
      sourceVault: {
        id: "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault",
      },
    },
    encryptionType: "EncryptionAtRestWithCustomerKey",
    identity: { type: "SystemAssigned" },
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

createADiskEncryptionSet().catch(console.error);
