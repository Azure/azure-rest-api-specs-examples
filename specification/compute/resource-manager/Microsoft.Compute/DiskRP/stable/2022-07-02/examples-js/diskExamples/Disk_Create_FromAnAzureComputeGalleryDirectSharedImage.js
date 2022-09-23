const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a disk.
 *
 * @summary Creates or updates a disk.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/diskExamples/Disk_Create_FromAnAzureComputeGalleryDirectSharedImage.json
 */
async function createAManagedDiskFromAnAzureComputeGalleryDirectSharedImage() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myDisk";
  const disk = {
    creationData: {
      createOption: "FromImage",
      galleryImageReference: {
        sharedGalleryImageId:
          "/SharedGalleries/{sharedGalleryUniqueName}/Images/{imageName}/Versions/1.0.0",
      },
    },
    location: "West US",
    osType: "Windows",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginCreateOrUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}

createAManagedDiskFromAnAzureComputeGalleryDirectSharedImage().catch(console.error);
