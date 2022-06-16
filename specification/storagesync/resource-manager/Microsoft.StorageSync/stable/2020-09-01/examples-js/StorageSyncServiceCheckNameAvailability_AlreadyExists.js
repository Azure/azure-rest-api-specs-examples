const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the give namespace name availability.
 *
 * @summary Check the give namespace name availability.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/StorageSyncServiceCheckNameAvailability_AlreadyExists.json
 */
async function storageSyncServiceCheckNameAvailabilityAlreadyExists() {
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

storageSyncServiceCheckNameAvailabilityAlreadyExists().catch(console.error);
