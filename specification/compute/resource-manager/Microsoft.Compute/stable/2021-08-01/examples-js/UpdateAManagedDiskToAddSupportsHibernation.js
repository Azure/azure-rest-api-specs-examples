const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateAManagedDiskToAddSupportsHibernation() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myDisk";
  const disk = { supportsHibernation: true };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}

updateAManagedDiskToAddSupportsHibernation().catch(console.error);
