const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a new queue with the specified queue name, under the specified account.
 *
 * @summary Creates a new queue with the specified queue name, under the specified account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/QueueOperationPutWithMetadata.json
 */
async function queueOperationPutWithMetadata() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res3376";
  const accountName = "sto328";
  const queueName = "queue6185";
  const queue = {
    metadata: { sample1: "meta1", sample2: "meta2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.queue.create(resourceGroupName, accountName, queueName, queue);
  console.log(result);
}
