const { ComputeManagementClient } = require("@azure/arm-compute-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an image.
 *
 * @summary Create or update an image.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-06-01/examples/CreateAnImageThatIncludesADataDiskFromAManagedDisk.json
 */
async function createAVirtualMachineImageThatIncludesADataDiskFromAManagedDisk() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const imageName = "myImage";
  const parameters = {
    location: "West US",
    storageProfile: {
      dataDisks: [
        {
          lun: 1,
          managedDisk: {
            id: "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk2",
          },
        },
      ],
      osDisk: {
        managedDisk: {
          id: "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk",
        },
        osState: "Generalized",
        osType: "Linux",
      },
      zoneResilient: false,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.images.beginCreateOrUpdateAndWait(
    resourceGroupName,
    imageName,
    parameters
  );
  console.log(result);
}
