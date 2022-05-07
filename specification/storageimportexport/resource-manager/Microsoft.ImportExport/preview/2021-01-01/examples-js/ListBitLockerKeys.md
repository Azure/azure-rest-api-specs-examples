Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storageimportexport_2.0.1/sdk/storageimportexport/arm-storageimportexport/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageImportExport } = require("@azure/arm-storageimportexport");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the BitLocker Keys for all drives in the specified job.
 *
 * @summary Returns the BitLocker Keys for all drives in the specified job.
 * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/ListBitLockerKeys.json
 */
async function listBitLockerKeysForDrivesInAJob() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const jobName = "myJob";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new StorageImportExport(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.bitLockerKeys.list(jobName, resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listBitLockerKeysForDrivesInAJob().catch(console.error);
```
