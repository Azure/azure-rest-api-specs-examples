const { ServiceFabricMeshManagementClient } = require("@azure/arm-servicefabricmesh");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the gateway resource identified by the name.
 *
 * @summary Deletes the gateway resource identified by the name.
 * x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/gateways/delete.json
 */
async function deleteGateway() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sbz_demo";
  const gatewayResourceName = "sampleGateway";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricMeshManagementClient(credential, subscriptionId);
  const result = await client.gateway.delete(resourceGroupName, gatewayResourceName);
  console.log(result);
}

deleteGateway().catch(console.error);
