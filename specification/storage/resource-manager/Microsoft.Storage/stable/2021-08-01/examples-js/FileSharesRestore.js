const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function restoreShares() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const shareName = "share1249";
  const deletedShare = {
    deletedShareName: "share1249",
    deletedShareVersion: "1234567890",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.restore(
    resourceGroupName,
    accountName,
    shareName,
    deletedShare
  );
  console.log(result);
}

restoreShares().catch(console.error);
