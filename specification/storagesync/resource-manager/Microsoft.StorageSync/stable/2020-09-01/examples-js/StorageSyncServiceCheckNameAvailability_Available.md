```javascript
const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the give namespace name availability.
 *
 * @summary Check the give namespace name availability.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/StorageSyncServiceCheckNameAvailability_Available.json
 */
async function storageSyncServiceCheckNameAvailabilityAvailable() {
  const subscriptionId = "5c6bc8e1-1eaf-4192-94d8-58ce463ac86c";
  const locationName = "westus";
  const parameters = {
    name: "newstoragesyncservicename",
    type: "Microsoft.StorageSync/storageSyncServices",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.storageSyncServices.checkNameAvailability(locationName, parameters);
  console.log(result);
}

storageSyncServiceCheckNameAvailabilityAvailable().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storagesync_9.0.1/sdk/storagesync/arm-storagesync/README.md) on how to add the SDK to your project and authenticate.
