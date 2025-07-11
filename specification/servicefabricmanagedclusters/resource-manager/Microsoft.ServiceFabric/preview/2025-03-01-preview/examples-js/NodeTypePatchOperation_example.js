const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update the configuration of a node type of a given managed cluster, only updating tags.
 *
 * @summary update the configuration of a node type of a given managed cluster, only updating tags.
 * x-ms-original-file: 2025-03-01-preview/NodeTypePatchOperation_example.json
 */
async function patchANodeType() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.nodeTypes.update("resRg", "myCluster", "BE", {
    tags: { a: "b" },
  });
  console.log(result);
}
