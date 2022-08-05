const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getInformationAboutADiskAccessResourceWithPrivateEndpoints() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskAccessName = "myDiskAccess";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskAccesses.get(resourceGroupName, diskAccessName);
  console.log(result);
}

getInformationAboutADiskAccessResourceWithPrivateEndpoints().catch(console.error);
