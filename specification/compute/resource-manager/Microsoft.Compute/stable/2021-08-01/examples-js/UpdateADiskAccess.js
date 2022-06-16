const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateADiskAccessResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskAccessName = "myDiskAccess";
  const diskAccess = {
    tags: { department: "Development", project: "PrivateEndpoints" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskAccesses.beginUpdateAndWait(
    resourceGroupName,
    diskAccessName,
    diskAccess
  );
  console.log(result);
}

updateADiskAccessResource().catch(console.error);
