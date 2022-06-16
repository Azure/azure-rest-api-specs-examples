const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified private endpoint by resource group.
 *
 * @summary Gets the specified private endpoint by resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PrivateEndpointGetForManualApproval.json
 */
async function getPrivateEndpointWithManualApprovalConnection() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const privateEndpointName = "testPe";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateEndpoints.get(resourceGroupName, privateEndpointName);
  console.log(result);
}

getPrivateEndpointWithManualApprovalConnection().catch(console.error);
