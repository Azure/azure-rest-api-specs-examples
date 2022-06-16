const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateContainers() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const containerName = "container6185";
  const blobContainer = {
    metadata: { metadata: "true" },
    publicAccess: "Container",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.update(
    resourceGroupName,
    accountName,
    containerName,
    blobContainer
  );
  console.log(result);
}

updateContainers().catch(console.error);
