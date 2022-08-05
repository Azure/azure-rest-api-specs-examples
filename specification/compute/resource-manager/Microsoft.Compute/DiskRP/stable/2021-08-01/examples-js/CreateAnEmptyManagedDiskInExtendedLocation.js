const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function createAnEmptyManagedDiskInExtendedLocation() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myDisk";
  const disk = {
    creationData: { createOption: "Empty" },
    diskSizeGB: 200,
    extendedLocation: { name: "{edge-zone-id}", type: "EdgeZone" },
    location: "West US",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginCreateOrUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}

createAnEmptyManagedDiskInExtendedLocation().catch(console.error);
