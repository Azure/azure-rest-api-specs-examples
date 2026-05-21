const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets properties of a specified share.
 *
 * @summary gets properties of a specified share.
 * x-ms-original-file: 2025-08-01/FileSharesGet_ProvisionedV2.json
 */
async function getShareProvisionedV2() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.get("res9871", "sto6217", "share1634");
  console.log(result);
}
