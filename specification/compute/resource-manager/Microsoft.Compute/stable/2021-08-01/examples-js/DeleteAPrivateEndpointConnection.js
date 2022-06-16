const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteAPrivateEndpointConnectionUnderADiskAccessResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskAccessName = "myDiskAccess";
  const privateEndpointConnectionName = "myPrivateEndpointConnection";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskAccesses.beginDeleteAPrivateEndpointConnectionAndWait(
    resourceGroupName,
    diskAccessName,
    privateEndpointConnectionName
  );
  console.log(result);
}

deleteAPrivateEndpointConnectionUnderADiskAccessResource().catch(console.error);
