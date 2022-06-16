const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function createAManagedDiskAndAssociateWithDiskAccessResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myDisk";
  const disk = {
    creationData: { createOption: "Empty" },
    diskAccessId:
      "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/{existing-diskAccess-name}",
    diskSizeGB: 200,
    location: "West US",
    networkAccessPolicy: "AllowPrivate",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginCreateOrUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}

createAManagedDiskAndAssociateWithDiskAccessResource().catch(console.error);
