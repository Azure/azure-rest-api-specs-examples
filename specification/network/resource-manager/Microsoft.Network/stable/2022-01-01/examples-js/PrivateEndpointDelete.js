const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified private endpoint.
 *
 * @summary Deletes the specified private endpoint.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/PrivateEndpointDelete.json
 */
async function deletePrivateEndpoint() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const privateEndpointName = "testPe";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateEndpoints.beginDeleteAndWait(
    resourceGroupName,
    privateEndpointName
  );
  console.log(result);
}

deletePrivateEndpoint().catch(console.error);
