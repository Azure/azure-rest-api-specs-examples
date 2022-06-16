const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function createAManagedUploadDisk() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myDisk";
  const disk = {
    creationData: { createOption: "Upload", uploadSizeBytes: 10737418752 },
    location: "West US",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginCreateOrUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}

createAManagedUploadDisk().catch(console.error);
