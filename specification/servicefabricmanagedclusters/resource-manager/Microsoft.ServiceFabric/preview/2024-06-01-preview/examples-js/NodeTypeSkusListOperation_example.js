const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Service Fabric node type supported SKUs.
 *
 * @summary Get a Service Fabric node type supported SKUs.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/examples/NodeTypeSkusListOperation_example.json
 */
async function listANodeTypeSkUs() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRICMANAGEDCLUSTERS_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const nodeTypeName = "BE";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.nodeTypeSkus.list(resourceGroupName, clusterName, nodeTypeName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
