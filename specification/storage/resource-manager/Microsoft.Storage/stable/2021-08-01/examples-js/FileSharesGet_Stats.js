const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function getShareStats() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9871";
  const accountName = "sto6217";
  const shareName = "share1634";
  const expand = "stats";
  const options = { expand: expand };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.get(resourceGroupName, accountName, shareName, options);
  console.log(result);
}

getShareStats().catch(console.error);
