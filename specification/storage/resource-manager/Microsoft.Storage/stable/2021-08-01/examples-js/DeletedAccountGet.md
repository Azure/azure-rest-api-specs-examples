Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function deletedAccountGet() {
  const subscriptionId = "{subscription-id}";
  const deletedAccountName = "sto1125";
  const location = "eastus";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.deletedAccounts.get(deletedAccountName, location);
  console.log(result);
}

deletedAccountGet().catch(console.error);
```
