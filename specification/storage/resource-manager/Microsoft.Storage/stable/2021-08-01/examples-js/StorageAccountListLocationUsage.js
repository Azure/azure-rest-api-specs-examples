const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function usageList() {
  const subscriptionId = "{subscription-id}";
  const location = "eastus2(stage)";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usages.listByLocation(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

usageList().catch(console.error);
