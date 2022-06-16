const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

async function deletesAPrivateEndpointConnectionWithAGivenName() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myResourceGroup";
  const scopeName = "myPrivateLinkScope";
  const privateEndpointConnectionName = "private-endpoint-connection-name";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginDeleteAndWait(
    resourceGroupName,
    scopeName,
    privateEndpointConnectionName
  );
  console.log(result);
}

deletesAPrivateEndpointConnectionWithAGivenName().catch(console.error);
