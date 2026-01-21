const { StorageClient } = require("@azure/arm-dell-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a FileSystemResource
 *
 * @summary get a FileSystemResource
 * x-ms-original-file: 2025-03-21/FileSystems_Get_MinimumSet_Gen.json
 */
async function fileSystemsGetMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "BF7E7352-2FE4-4163-9CF7-5FF8EC2E9B92";
  const client = new StorageClient(credential, subscriptionId);
  const result = await client.fileSystems.get("rgDell", "abcd");
  console.log(result);
}
