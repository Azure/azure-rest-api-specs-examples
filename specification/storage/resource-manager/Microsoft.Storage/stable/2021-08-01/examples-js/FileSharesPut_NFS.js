const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function createNfsShares() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res346";
  const accountName = "sto666";
  const shareName = "share1235";
  const fileShare = { enabledProtocols: "NFS" };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.create(
    resourceGroupName,
    accountName,
    shareName,
    fileShare
  );
  console.log(result);
}

createNfsShares().catch(console.error);
