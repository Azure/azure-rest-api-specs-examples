const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Approve or reject private end point connection for a private link service in a subscription.
 *
 * @summary Approve or reject private end point connection for a private link service in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PrivateLinkServiceUpdatePrivateEndpointConnection.json
 */
async function approveOrRejectPrivateEndPointConnectionForAPrivateLinkService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const serviceName = "testPls";
  const peConnectionName = "testPlePeConnection";
  const parameters = {
    name: "testPlePeConnection",
    privateEndpoint: {
      id: "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe",
    },
    privateLinkServiceConnectionState: {
      description: "approved it for some reason.",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateLinkServices.updatePrivateEndpointConnection(
    resourceGroupName,
    serviceName,
    peConnectionName,
    parameters
  );
  console.log(result);
}

approveOrRejectPrivateEndPointConnectionForAPrivateLinkService().catch(console.error);
