Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Abort live Migration of storage account to enable Hns
 *
 * @summary Abort live Migration of storage account to enable Hns
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountAbortHierarchicalNamespaceMigration.json
 */
async function storageAccountAbortHierarchicalNamespaceMigration() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4228";
  const accountName = "sto2434";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.beginAbortHierarchicalNamespaceMigrationAndWait(
    resourceGroupName,
    accountName
  );
  console.log(result);
}

storageAccountAbortHierarchicalNamespaceMigration().catch(console.error);
```
