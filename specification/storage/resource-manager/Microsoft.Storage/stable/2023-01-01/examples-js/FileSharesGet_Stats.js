const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets properties of a specified share.
 *
 * @summary Gets properties of a specified share.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/FileSharesGet_Stats.json
 */
async function getShareStats() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res9871";
  const accountName = "sto6217";
  const shareName = "share1634";
  const expand = "stats";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.get(resourceGroupName, accountName, shareName, options);
  console.log(result);
}
