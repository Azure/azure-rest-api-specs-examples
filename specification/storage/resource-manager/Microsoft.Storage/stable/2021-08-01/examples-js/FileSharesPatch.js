const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateShares() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const shareName = "share6185";
  const fileShare = { metadata: { type: "image" } };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.update(
    resourceGroupName,
    accountName,
    shareName,
    fileShare
  );
  console.log(result);
}

updateShares().catch(console.error);
