const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Service Fabric cluster resource created or in the process of being created in the specified resource group.
 *
 * @summary Get a Service Fabric cluster resource created or in the process of being created in the specified resource group.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterGetOperation_example.json
 */
async function getACluster() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.clusters.get(resourceGroupName, clusterName);
  console.log(result);
}

getACluster().catch(console.error);
