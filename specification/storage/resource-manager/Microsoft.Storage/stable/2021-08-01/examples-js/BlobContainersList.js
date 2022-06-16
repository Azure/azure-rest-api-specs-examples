const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function listContainers() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9290";
  const accountName = "sto1590";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.blobContainers.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listContainers().catch(console.error);
