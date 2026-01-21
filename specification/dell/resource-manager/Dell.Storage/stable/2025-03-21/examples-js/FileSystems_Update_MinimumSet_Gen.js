const { StorageClient } = require("@azure/arm-dell-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a FileSystemResource
 *
 * @summary update a FileSystemResource
 * x-ms-original-file: 2025-03-21/FileSystems_Update_MinimumSet_Gen.json
 */
async function fileSystemsUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "BF7E7352-2FE4-4163-9CF7-5FF8EC2E9B92";
  const client = new StorageClient(credential, subscriptionId);
  const result = await client.fileSystems.update("rgDell", "abcd", {
    properties: { delegatedSubnetId: "uqfvajvyltgmqvdnxhbrfqbpuey", capacity: { current: "5" } },
  });
  console.log(result);
}
