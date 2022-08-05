const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function createAManagedDiskFromUploadPreparedSecureCreateOption() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myDisk";
  const disk = {
    creationData: {
      createOption: "UploadPreparedSecure",
      uploadSizeBytes: 10737418752,
    },
    location: "West US",
    osType: "Windows",
    securityProfile: { securityType: "TrustedLaunch" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginCreateOrUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}

createAManagedDiskFromUploadPreparedSecureCreateOption().catch(console.error);
