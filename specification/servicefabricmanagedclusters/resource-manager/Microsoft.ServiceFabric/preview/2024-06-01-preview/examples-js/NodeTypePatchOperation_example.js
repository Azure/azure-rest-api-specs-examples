const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the configuration of a node type of a given managed cluster, only updating tags.
 *
 * @summary Update the configuration of a node type of a given managed cluster, only updating tags.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/examples/NodeTypePatchOperation_example.json
 */
async function patchANodeType() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRICMANAGEDCLUSTERS_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const nodeTypeName = "BE";
  const parameters = { tags: { a: "b" } };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.nodeTypes.update(
    resourceGroupName,
    clusterName,
    nodeTypeName,
    parameters,
  );
  console.log(result);
}
