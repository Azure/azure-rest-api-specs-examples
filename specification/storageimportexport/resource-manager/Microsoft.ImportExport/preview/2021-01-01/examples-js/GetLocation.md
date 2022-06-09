```javascript
const { StorageImportExport } = require("@azure/arm-storageimportexport");
const { DefaultAzureCredential } = require("@azure/identity");

async function getLocations() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const locationName = "West US";
  const credential = new DefaultAzureCredential();
  const client = new StorageImportExport(credential, subscriptionId);
  const result = await client.locations.get(locationName);
  console.log(result);
}

getLocations().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storageimportexport_2.0.1/sdk/storageimportexport/arm-storageimportexport/README.md) on how to add the SDK to your project and authenticate.
