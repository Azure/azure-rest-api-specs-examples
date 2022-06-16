const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function deletedAccountList() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.deletedAccounts.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

deletedAccountList().catch(console.error);
