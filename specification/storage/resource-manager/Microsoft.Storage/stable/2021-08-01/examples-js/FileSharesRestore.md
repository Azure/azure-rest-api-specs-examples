Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function restoreShares() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const shareName = "share1249";
  const deletedShare = {
    deletedShareName: "share1249",
    deletedShareVersion: "1234567890",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.restore(
    resourceGroupName,
    accountName,
    shareName,
    deletedShare
  );
  console.log(result);
}

restoreShares().catch(console.error);
```
