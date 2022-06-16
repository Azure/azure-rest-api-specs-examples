const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function setLegalHoldContainersWithAllowProtectedAppendWritesAll() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4303";
  const accountName = "sto7280";
  const containerName = "container8723";
  const legalHold = {
    allowProtectedAppendWritesAll: true,
    tags: ["tag1", "tag2", "tag3"],
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.setLegalHold(
    resourceGroupName,
    accountName,
    containerName,
    legalHold
  );
  console.log(result);
}

setLegalHoldContainersWithAllowProtectedAppendWritesAll().catch(console.error);
