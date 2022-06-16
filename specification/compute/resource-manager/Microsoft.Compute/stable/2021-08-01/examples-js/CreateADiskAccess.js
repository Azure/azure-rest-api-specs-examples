const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function createADiskAccessResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskAccessName = "myDiskAccess";
  const diskAccess = { location: "West US" };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskAccesses.beginCreateOrUpdateAndWait(
    resourceGroupName,
    diskAccessName,
    diskAccess
  );
  console.log(result);
}

createADiskAccessResource().catch(console.error);
