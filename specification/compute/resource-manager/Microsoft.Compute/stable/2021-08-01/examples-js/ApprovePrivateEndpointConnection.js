const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function approveAPrivateEndpointConnectionUnderADiskAccessResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskAccessName = "myDiskAccess";
  const privateEndpointConnectionName = "myPrivateEndpointConnection";
  const privateEndpointConnection = {
    privateLinkServiceConnectionState: {
      description: "Approving myPrivateEndpointConnection",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskAccesses.beginUpdateAPrivateEndpointConnectionAndWait(
    resourceGroupName,
    diskAccessName,
    privateEndpointConnectionName,
    privateEndpointConnection
  );
  console.log(result);
}

approveAPrivateEndpointConnectionUnderADiskAccessResource().catch(console.error);
