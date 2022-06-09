```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Live Migration of storage account to enable Hns
 *
 * @summary Live Migration of storage account to enable Hns
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountHierarchicalNamespaceMigration.json
 */
async function storageAccountHierarchicalNamespaceMigration() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4228";
  const accountName = "sto2434";
  const requestType = "HnsOnValidationRequest";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.beginHierarchicalNamespaceMigrationAndWait(
    resourceGroupName,
    accountName,
    requestType
  );
  console.log(result);
}

storageAccountHierarchicalNamespaceMigration().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.
