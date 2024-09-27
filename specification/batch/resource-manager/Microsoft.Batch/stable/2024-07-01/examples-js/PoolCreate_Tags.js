const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new pool inside the specified account.
 *
 * @summary Creates a new pool inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/PoolCreate_Tags.json
 */
async function createPoolTags() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["BATCH_RESOURCE_GROUP"] || "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const parameters = {
    deploymentConfiguration: {
      virtualMachineConfiguration: {
        imageReference: {
          offer: "0001-com-ubuntu-server-jammy",
          publisher: "Canonical",
          sku: "22_04-lts",
          version: "latest",
        },
        nodeAgentSkuId: "batch.node.ubuntu 22.04",
      },
    },
    scaleSettings: {
      fixedScale: { targetDedicatedNodes: 1, targetLowPriorityNodes: 0 },
    },
    tags: { tagName1: "TagValue1", tagName2: "TagValue2" },
    vmSize: "Standard_d4s_v3",
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.poolOperations.create(
    resourceGroupName,
    accountName,
    poolName,
    parameters,
  );
  console.log(result);
}
