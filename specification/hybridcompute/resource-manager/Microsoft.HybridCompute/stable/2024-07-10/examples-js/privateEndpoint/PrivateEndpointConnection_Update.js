const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Approve or reject a private endpoint connection with a given name.
 *
 * @summary Approve or reject a private endpoint connection with a given name.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/privateEndpoint/PrivateEndpointConnection_Update.json
 */
async function approveOrRejectAPrivateEndpointConnectionWithAGivenName() {
  const subscriptionId =
    process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const scopeName = "myPrivateLinkScope";
  const privateEndpointConnectionName = "private-endpoint-connection-name";
  const parameters = {
    properties: {
      privateLinkServiceConnectionState: {
        description: "Approved by johndoe@contoso.com",
        status: "Approved",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    scopeName,
    privateEndpointConnectionName,
    parameters,
  );
  console.log(result);
}
