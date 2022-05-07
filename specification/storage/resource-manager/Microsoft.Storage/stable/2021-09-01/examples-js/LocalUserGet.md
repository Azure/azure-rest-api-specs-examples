Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the local user of the storage account by username.
 *
 * @summary Get the local user of the storage account by username.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/LocalUserGet.json
 */
async function getLocalUser() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const accountName = "sto2527";
  const username = "user1";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.localUsersOperations.get(resourceGroupName, accountName, username);
  console.log(result);
}

getLocalUser().catch(console.error);
```
