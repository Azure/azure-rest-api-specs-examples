const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the information about all services of an application resource. The information include the description and other properties of the Service.
 *
 * @summary Gets the information about all services of an application resource. The information include the description and other properties of the Service.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/applications/services/list.json
 */
async function listServices() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sbz_demo";
  const applicationResourceName = "sampleApplication";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.service.list(resourceGroupName, applicationResourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listServices().catch(console.error);
