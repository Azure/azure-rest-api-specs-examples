const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function queueOperationPutWithMetadata() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
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

queueOperationPutWithMetadata().catch(console.error);
