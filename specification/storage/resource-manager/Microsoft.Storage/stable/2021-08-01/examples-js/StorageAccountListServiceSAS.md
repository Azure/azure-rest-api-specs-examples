Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountListServiceSas() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7439";
  const accountName = "sto1299";
  const parameters = {
    canonicalizedResource: "/blob/sto1299/music",
    sharedAccessExpiryTime: new Date("2017-05-24T11:32:48.8457197Z"),
    permissions: "l",
    resource: "c",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.listServiceSAS(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

storageAccountListServiceSas().catch(console.error);
```
