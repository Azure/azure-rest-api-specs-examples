const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteImmutabilityPolicy() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res1581";
  const accountName = "sto9621";
  const containerName = "container4910";
  const ifMatch = '"8d59f81a7fa7be0"';
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.deleteImmutabilityPolicy(
    resourceGroupName,
    accountName,
    containerName,
    ifMatch
  );
  console.log(result);
}

deleteImmutabilityPolicy().catch(console.error);
