const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the list of available Service Fabric resource provider API operations.
 *
 * @summary Get the list of available Service Fabric resource provider API operations.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/Operations_example.json
 */
async function listAvailableOperations() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
