const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteContainers() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4079";
  const accountName = "sto4506";
  const containerName = "container9689";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.delete(resourceGroupName, accountName, containerName);
  console.log(result);
}

deleteContainers().catch(console.error);
