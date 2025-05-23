const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a disk.
 *
 * @summary Creates or updates a disk.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskExamples/Disk_Create_ConfidentialVMSupportedDiskEncryptedWithCMK.json
 */
async function createAConfidentialVMSupportedDiskEncryptedWithCustomerManagedKey() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const diskName = "myDisk";
  const disk = {
    creationData: {
      createOption: "FromImage",
      imageReference: {
        id: "/Subscriptions/{subscriptionId}/Providers/Microsoft.Compute/Locations/westus/Publishers/{publisher}/ArtifactTypes/VMImage/Offers/{offer}/Skus/{sku}/Versions/1.0.0",
      },
    },
    location: "West US",
    osType: "Windows",
    securityProfile: {
      secureVMDiskEncryptionSetId:
        "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}",
      securityType: "ConfidentialVM_DiskEncryptedWithCustomerKey",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginCreateOrUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}
