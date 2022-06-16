const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function listAllPossiblePrivateLinkResourcesUnderDiskAccessResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskAccessName = "myDiskAccess";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskAccesses.getPrivateLinkResources(
    resourceGroupName,
    diskAccessName
  );
  console.log(result);
}

listAllPossiblePrivateLinkResourcesUnderDiskAccessResource().catch(console.error);
