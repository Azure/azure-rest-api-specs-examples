const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all Service Fabric cluster resources created or in the process of being created in the subscription.
 *
 * @summary Gets all Service Fabric cluster resources created or in the process of being created in the subscription.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterListOperation_example.json
 */
async function listClusters() {
  const subscriptionId =
    process.env["SERVICEFABRIC_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusters.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
