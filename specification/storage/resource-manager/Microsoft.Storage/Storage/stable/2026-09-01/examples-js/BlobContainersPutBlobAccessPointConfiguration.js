const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new container under the specified account as described by request body. The container resource includes metadata and properties for that container. It does not include a list of the blobs contained by the container.
 *
 * @summary creates a new container under the specified account as described by request body. The container resource includes metadata and properties for that container. It does not include a list of the blobs contained by the container.
 * x-ms-original-file: 2026-09-01/BlobContainersPutBlobAccessPointConfiguration.json
 */
async function putContainerWithBlobAccessPointConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.create("res3376", "sto328", "container6185", {
    blobAccessPointConfiguration: {
      blobAccessPointConfigurationName: "myAccessPointConfig",
      blobAccessPointConfigurationUniqueId: "00000000-0000-0000-0000-000000000000",
    },
  });
  console.log(result);
}
