```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the local user associated with the specified storage account.
 *
 * @summary Deletes the local user associated with the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/LocalUserDelete.json
 */
async function deleteLocalUser() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const accountName = "sto2527";
  const username = "user1";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.localUsersOperations.delete(resourceGroupName, accountName, username);
  console.log(result);
}

deleteLocalUser().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.1/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.
