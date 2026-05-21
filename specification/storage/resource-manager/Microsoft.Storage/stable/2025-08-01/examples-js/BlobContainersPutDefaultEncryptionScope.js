const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new container under the specified account as described by request body. The container resource includes metadata and properties for that container. It does not include a list of the blobs contained by the container.
 *
 * @summary creates a new container under the specified account as described by request body. The container resource includes metadata and properties for that container. It does not include a list of the blobs contained by the container.
 * x-ms-original-file: 2025-08-01/BlobContainersPutDefaultEncryptionScope.json
 */
async function putContainerWithDefaultEncryptionScope() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.create("res3376", "sto328", "container6185", {
    defaultEncryptionScope: "encryptionscope185",
    denyEncryptionScopeOverride: true,
  });
  console.log(result);
}
