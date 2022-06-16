const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function listDeletedContainers() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9290";
  const accountName = "sto1590";
  const include = "deleted";
  const options = { include: include };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.blobContainers.list(resourceGroupName, accountName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDeletedContainers().catch(console.error);
