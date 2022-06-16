const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function listShareSnapshots() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9290";
  const accountName = "sto1590";
  const expand = "snapshots";
  const options = { expand: expand };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fileShares.list(resourceGroupName, accountName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listShareSnapshots().catch(console.error);
