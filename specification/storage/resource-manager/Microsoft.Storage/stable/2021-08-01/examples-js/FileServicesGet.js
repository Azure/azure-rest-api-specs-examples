const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function getFileServices() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4410";
  const accountName = "sto8607";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileServices.getServiceProperties(resourceGroupName, accountName);
  console.log(result);
}

getFileServices().catch(console.error);
