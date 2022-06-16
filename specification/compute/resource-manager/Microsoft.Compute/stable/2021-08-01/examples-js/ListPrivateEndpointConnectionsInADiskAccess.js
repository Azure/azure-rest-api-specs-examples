const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getInformationAboutAPrivateEndpointConnectionUnderADiskAccessResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskAccessName = "myDiskAccess";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.diskAccesses.listPrivateEndpointConnections(
    resourceGroupName,
    diskAccessName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getInformationAboutAPrivateEndpointConnectionUnderADiskAccessResource().catch(console.error);
