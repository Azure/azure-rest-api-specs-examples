const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the properties of a storage account’s Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
 *
 * @summary Sets the properties of a storage account’s Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/BlobServicesPutLastAccessTimeBasedTracking.json
 */
async function blobServicesPutLastAccessTimeBasedTracking() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4410";
  const accountName = "sto8607";
  const parameters = {
    lastAccessTimeTrackingPolicy: {
      name: "AccessTimeTracking",
      blobType: ["blockBlob"],
      enable: true,
      trackingGranularityInDays: 1,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobServices.setServiceProperties(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

blobServicesPutLastAccessTimeBasedTracking().catch(console.error);
