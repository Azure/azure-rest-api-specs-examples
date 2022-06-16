const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getsListOfPrivateEndpointConnectionsOnAPrivateLinkScope() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myResourceGroup";
  const scopeName = "myPrivateLinkScope";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.listByPrivateLinkScope(
    resourceGroupName,
    scopeName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsListOfPrivateEndpointConnectionsOnAPrivateLinkScope().catch(console.error);
