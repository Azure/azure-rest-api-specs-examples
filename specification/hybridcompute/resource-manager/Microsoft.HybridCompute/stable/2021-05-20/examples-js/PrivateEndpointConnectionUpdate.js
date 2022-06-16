const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

async function approveOrRejectAPrivateEndpointConnectionWithAGivenName() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myResourceGroup";
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
    parameters
  );
  console.log(result);
}

approveOrRejectAPrivateEndpointConnectionWithAGivenName().catch(console.error);
