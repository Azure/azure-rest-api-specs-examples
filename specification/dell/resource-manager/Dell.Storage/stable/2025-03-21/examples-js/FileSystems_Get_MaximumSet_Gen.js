const { StorageClient } = require("@azure/arm-dell-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a FileSystemResource
 *
 * @summary get a FileSystemResource
 * x-ms-original-file: 2025-03-21/FileSystems_Get_MaximumSet_Gen.json
 */
async function fileSystemsGetMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "4B6E265D-57CF-4A9D-8B35-3CC68ED9D208";
  const client = new StorageClient(credential, subscriptionId);
  const result = await client.fileSystems.get("rgDell", "abcd");
  console.log(result);
}
