const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function queueServicesList() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9290";
  const accountName = "sto1590";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.queueServices.list(resourceGroupName, accountName);
  console.log(result);
}

queueServicesList().catch(console.error);
