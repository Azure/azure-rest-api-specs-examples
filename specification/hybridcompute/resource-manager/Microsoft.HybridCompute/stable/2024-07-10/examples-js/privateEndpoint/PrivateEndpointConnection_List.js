const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all private endpoint connections on a private link scope.
 *
 * @summary Gets all private endpoint connections on a private link scope.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/privateEndpoint/PrivateEndpointConnection_List.json
 */
async function getsListOfPrivateEndpointConnectionsOnAPrivateLinkScope() {
  const subscriptionId =
    process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const scopeName = "myPrivateLinkScope";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.listByPrivateLinkScope(
    resourceGroupName,
    scopeName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
