const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a disk.
 *
 * @summary creates or updates a disk.
 * x-ms-original-file: 2025-01-02/diskExamples/Disk_Create_WithSecurityProfile.json
 */
async function createAManagedDiskWithSecurityProfile() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.createOrUpdate("myResourceGroup", "myDisk", {
    location: "North Central US",
    osType: "Windows",
    securityProfile: { securityType: "TrustedLaunch" },
    creationData: {
      createOption: "FromImage",
      imageReference: {
        id: "/Subscriptions/{subscriptionId}/Providers/Microsoft.Compute/Locations/uswest/Publishers/Microsoft/ArtifactTypes/VMImage/Offers/{offer}",
      },
    },
  });
  console.log(result);
}
