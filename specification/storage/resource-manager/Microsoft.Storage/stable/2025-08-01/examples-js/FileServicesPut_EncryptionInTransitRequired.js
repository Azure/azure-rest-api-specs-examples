const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to sets the properties of file services in storage accounts, including CORS (Cross-Origin Resource Sharing) rules.
 *
 * @summary sets the properties of file services in storage accounts, including CORS (Cross-Origin Resource Sharing) rules.
 * x-ms-original-file: 2025-08-01/FileServicesPut_EncryptionInTransitRequired.json
 */
async function putFileServicesEncryptionInTransitRequired() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileServices.setServiceProperties("res4410", "sto8607", {
    protocolSettings: {
      nfs: { encryptionInTransit: { required: true } },
      smb: { encryptionInTransit: { required: true } },
    },
  });
  console.log(result);
}
