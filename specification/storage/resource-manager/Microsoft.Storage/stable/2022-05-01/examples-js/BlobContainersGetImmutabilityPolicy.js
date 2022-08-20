const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the existing immutability policy along with the corresponding ETag in response headers and body.
 *
 * @summary Gets the existing immutability policy along with the corresponding ETag in response headers and body.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/BlobContainersGetImmutabilityPolicy.json
 */
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
