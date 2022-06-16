const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about an available Service Fabric cluster code version.
 *
 * @summary Gets information about an available Service Fabric cluster code version.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterVersionsGet_example.json
 */
async function getClusterVersion() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const clusterVersion = "6.1.480.9494";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.clusterVersions.get(location, clusterVersion);
  console.log(result);
}

getClusterVersion().catch(console.error);
