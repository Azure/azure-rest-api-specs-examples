const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function getImmutabilityPolicy() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res5221";
  const accountName = "sto9177";
  const containerName = "container3489";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.getImmutabilityPolicy(
    resourceGroupName,
    accountName,
    containerName
  );
  console.log(result);
}

getImmutabilityPolicy().catch(console.error);
