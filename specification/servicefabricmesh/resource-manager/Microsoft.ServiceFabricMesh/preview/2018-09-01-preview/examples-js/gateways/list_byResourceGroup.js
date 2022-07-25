const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the information about all gateway resources in a given resource group. The information include the description and other properties of the Gateway.
 *
 * @summary Gets the information about all gateway resources in a given resource group. The information include the description and other properties of the Gateway.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/gateways/list_byResourceGroup.json
 */
async function listGatewaysByResourceGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sbz_demo";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.gateway.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listGatewaysByResourceGroup().catch(console.error);
