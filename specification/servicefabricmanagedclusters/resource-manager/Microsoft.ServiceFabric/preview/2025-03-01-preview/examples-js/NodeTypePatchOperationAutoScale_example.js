const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update the configuration of a node type of a given managed cluster, only updating tags.
 *
 * @summary update the configuration of a node type of a given managed cluster, only updating tags.
 * x-ms-original-file: 2025-03-01-preview/NodeTypePatchOperationAutoScale_example.json
 */
async function patchANodeTypeWhileAutoScaling() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.nodeTypes.update("resRg", "myCluster", "BE", {
    sku: { name: "Standard_S0", capacity: 10, tier: "Standard" },
    tags: { a: "b" },
  });
  console.log(result);
}
