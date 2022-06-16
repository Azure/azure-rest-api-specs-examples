const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteShares() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4079";
  const accountName = "sto4506";
  const shareName = "share9689";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.delete(resourceGroupName, accountName, shareName);
  console.log(result);
}

deleteShares().catch(console.error);
