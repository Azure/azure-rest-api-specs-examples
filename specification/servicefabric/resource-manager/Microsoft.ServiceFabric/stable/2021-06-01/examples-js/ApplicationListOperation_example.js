const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all application resources created or in the process of being created in the Service Fabric cluster resource.
 *
 * @summary Gets all application resources created or in the process of being created in the Service Fabric cluster resource.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationListOperation_example.json
 */
async function getAListOfApplicationResources() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.applications.list(resourceGroupName, clusterName);
  console.log(result);
}

getAListOfApplicationResources().catch(console.error);
