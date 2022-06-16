const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function getBlobContainersGetWithAllowProtectedAppendWritesAll() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9871";
  const accountName = "sto6217";
  const containerName = "container1634";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.get(resourceGroupName, accountName, containerName);
  console.log(result);
}

getBlobContainersGetWithAllowProtectedAppendWritesAll().catch(console.error);
