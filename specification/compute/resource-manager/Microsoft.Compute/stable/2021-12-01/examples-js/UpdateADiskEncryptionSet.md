Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates (patches) a disk encryption set.
 *
 * @summary Updates (patches) a disk encryption set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/UpdateADiskEncryptionSet.json
 */
async function updateADiskEncryptionSet() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskEncryptionSetName = "myDiskEncryptionSet";
  const diskEncryptionSet = {
    activeKey: {
      keyUrl: "https://myvmvault.vault-int.azure-int.net/keys/keyName/keyVersion",
      sourceVault: {
        id: "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault",
      },
    },
    encryptionType: "EncryptionAtRestWithCustomerKey",
    tags: { department: "Development", project: "Encryption" },
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

updateADiskEncryptionSet().catch(console.error);
```
