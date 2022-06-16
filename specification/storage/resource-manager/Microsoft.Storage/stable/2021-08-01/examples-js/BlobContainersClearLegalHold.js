const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function clearLegalHoldContainers() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4303";
  const accountName = "sto7280";
  const containerName = "container8723";
  const legalHold = { tags: ["tag1", "tag2", "tag3"] };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.clearLegalHold(
    resourceGroupName,
    accountName,
    containerName,
    legalHold
  );
  console.log(result);
}

clearLegalHoldContainers().catch(console.error);
