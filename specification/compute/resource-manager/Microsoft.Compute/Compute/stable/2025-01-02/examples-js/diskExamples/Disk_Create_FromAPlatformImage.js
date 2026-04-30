const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a disk.
 *
 * @summary creates or updates a disk.
 * x-ms-original-file: 2025-01-02/diskExamples/Disk_Create_FromAPlatformImage.json
 */
async function createAManagedDiskFromAPlatformImage() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscriptionId}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.createOrUpdate("myResourceGroup", "myDisk", {
    location: "West US",
    osType: "Windows",
    creationData: {
      createOption: "FromImage",
      imageReference: {
        id: "/Subscriptions/{subscriptionId}/Providers/Microsoft.Compute/Locations/westus/Publishers/{publisher}/ArtifactTypes/VMImage/Offers/{offer}/Skus/{sku}/Versions/1.0.0",
      },
    },
  });
  console.log(result);
}
