Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the storage account name is valid and is not already in use.
 *
 * @summary Checks that the storage account name is valid and is not already in use.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountCheckNameAvailability.json
 */
async function storageAccountCheckNameAvailability() {
  const subscriptionId = "{subscription-id}";
  const accountName = {
    name: "sto3363",
    type: "Microsoft.Storage/storageAccounts",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.checkNameAvailability(accountName);
  console.log(result);
}

storageAccountCheckNameAvailability().catch(console.error);
```
