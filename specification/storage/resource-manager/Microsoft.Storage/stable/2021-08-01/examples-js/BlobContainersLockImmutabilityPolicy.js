const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function lockImmutabilityPolicy() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res2702";
  const accountName = "sto5009";
  const containerName = "container1631";
  const ifMatch = '"8d59f825b721dd3"';
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.lockImmutabilityPolicy(
    resourceGroupName,
    accountName,
    containerName,
    ifMatch
  );
  console.log(result);
}

lockImmutabilityPolicy().catch(console.error);
