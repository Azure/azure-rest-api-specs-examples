const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates (patches) a disk access resource.
 *
 * @summary Updates (patches) a disk access resource.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/diskAccessExamples/DiskAccess_Update.json
 */
async function updateADiskAccessResource() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const diskAccessName = "myDiskAccess";
  const diskAccess = {
    tags: { department: "Development", project: "PrivateEndpoints" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskAccesses.beginUpdateAndWait(
    resourceGroupName,
    diskAccessName,
    diskAccess,
  );
  console.log(result);
}
