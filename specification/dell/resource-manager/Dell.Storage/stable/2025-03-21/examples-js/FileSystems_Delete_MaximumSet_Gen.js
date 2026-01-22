const { StorageClient } = require("@azure/arm-dell-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a FileSystemResource
 *
 * @summary delete a FileSystemResource
 * x-ms-original-file: 2025-03-21/FileSystems_Delete_MaximumSet_Gen.json
 */
async function fileSystemsDeleteMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "4B6E265D-57CF-4A9D-8B35-3CC68ED9D208";
  const client = new StorageClient(credential, subscriptionId);
  await client.fileSystems.delete("rgDell", "abcd");
}
