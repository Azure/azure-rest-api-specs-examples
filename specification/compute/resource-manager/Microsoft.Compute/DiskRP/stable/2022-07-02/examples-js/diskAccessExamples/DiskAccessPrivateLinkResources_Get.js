const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the private link resources possible under disk access resource
 *
 * @summary Gets the private link resources possible under disk access resource
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/diskAccessExamples/DiskAccessPrivateLinkResources_Get.json
 */
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
