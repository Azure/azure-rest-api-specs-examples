const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all application type name resources created or in the process of being created in the Service Fabric cluster resource.
 *
 * @summary Gets all application type name resources created or in the process of being created in the Service Fabric cluster resource.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationTypeNameListOperation_example.json
 */
async function getAListOfApplicationTypeNameResources() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.applicationTypes.list(resourceGroupName, clusterName);
  console.log(result);
}

getAListOfApplicationTypeNameResources().catch(console.error);
