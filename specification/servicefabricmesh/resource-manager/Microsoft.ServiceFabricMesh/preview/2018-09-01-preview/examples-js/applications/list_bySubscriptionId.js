const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the information about all application resources in a given resource group. The information include the description and other properties of the application.
 *
 * @summary Gets the information about all application resources in a given resource group. The information include the description and other properties of the application.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/applications/list_bySubscriptionId.json
 */
async function listApplicationsBySubscriptionId() {
  const subscriptionId =
    process.env["SERVICEFABRICMESH_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.application.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
