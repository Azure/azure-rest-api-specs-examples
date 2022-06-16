const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

async function signalRPrivateEndpointConnectionsUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const privateEndpointConnectionName = "mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "mySignalRService";
  const parameters = {
    privateEndpoint: {
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint",
    },
    privateLinkServiceConnectionState: {
      actionsRequired: "None",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalRPrivateEndpointConnections.update(
    privateEndpointConnectionName,
    resourceGroupName,
    resourceName,
    parameters
  );
  console.log(result);
}

signalRPrivateEndpointConnectionsUpdate().catch(console.error);
