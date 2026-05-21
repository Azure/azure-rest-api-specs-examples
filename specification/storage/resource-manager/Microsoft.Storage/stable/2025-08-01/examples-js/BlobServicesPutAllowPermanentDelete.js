const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to sets the properties of a storage account’s Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
 *
 * @summary sets the properties of a storage account’s Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
 * x-ms-original-file: 2025-08-01/BlobServicesPutAllowPermanentDelete.json
 */
async function blobServicesPutAllowPermanentDelete() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobServices.setServiceProperties("res4410", "sto8607", {
    deleteRetentionPolicy: { allowPermanentDelete: true, days: 300, enabled: true },
    isVersioningEnabled: true,
  });
  console.log(result);
}
