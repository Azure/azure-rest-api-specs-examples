const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates share properties as specified in request body. Properties not mentioned in the request will not be changed. Update fails if the specified share does not already exist.
 *
 * @summary Updates share properties as specified in request body. Properties not mentioned in the request will not be changed. Update fails if the specified share does not already exist.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/FileSharesPatch_PaidBursting.json
 */
async function updateSharePaidBursting() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res3376";
  const accountName = "sto328";
  const shareName = "share6185";
  const fileShare = {
    fileSharePaidBursting: {
      paidBurstingEnabled: true,
      paidBurstingMaxBandwidthMibps: 10340,
      paidBurstingMaxIops: 102400,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.update(
    resourceGroupName,
    accountName,
    shareName,
    fileShare,
  );
  console.log(result);
}
