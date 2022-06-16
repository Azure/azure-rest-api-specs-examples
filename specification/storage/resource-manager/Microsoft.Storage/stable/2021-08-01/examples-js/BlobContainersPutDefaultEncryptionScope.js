const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function putContainerWithDefaultEncryptionScope() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const containerName = "container6185";
  const blobContainer = {
    defaultEncryptionScope: "encryptionscope185",
    denyEncryptionScopeOverride: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.create(
    resourceGroupName,
    accountName,
    containerName,
    blobContainer
  );
  console.log(result);
}

putContainerWithDefaultEncryptionScope().catch(console.error);
