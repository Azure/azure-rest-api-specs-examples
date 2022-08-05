const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function createAnUltraManagedDiskWithLogicalSectorSize512E() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myDisk";
  const disk = {
    creationData: { createOption: "Empty", logicalSectorSize: 512 },
    diskSizeGB: 200,
    location: "West US",
    sku: { name: "UltraSSD_LRS" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginCreateOrUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}

createAnUltraManagedDiskWithLogicalSectorSize512E().catch(console.error);
