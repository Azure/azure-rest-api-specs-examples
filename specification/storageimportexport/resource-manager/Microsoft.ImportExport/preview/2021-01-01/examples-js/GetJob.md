Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storageimportexport_2.0.1/sdk/storageimportexport/arm-storageimportexport/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageImportExport } = require("@azure/arm-storageimportexport");
const { DefaultAzureCredential } = require("@azure/identity");

async function getImportJob() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const jobName = "myJob";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new StorageImportExport(credential, subscriptionId);
  const result = await client.jobs.get(jobName, resourceGroupName);
  console.log(result);
}

getImportJob().catch(console.error);
```
