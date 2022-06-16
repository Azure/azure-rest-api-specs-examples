const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function listLocalUsers() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const accountName = "sto2527";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.localUsersOperations.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listLocalUsers().catch(console.error);
