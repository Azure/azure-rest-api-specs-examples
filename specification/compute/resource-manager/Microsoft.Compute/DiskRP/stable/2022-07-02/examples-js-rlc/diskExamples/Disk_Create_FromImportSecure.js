const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { getLongRunningPoller } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Creates or updates a disk.
 *
 * @summary Creates or updates a disk.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/diskExamples/Disk_Create_FromImportSecure.json
 */
async function createAManagedDiskFromImportSecureCreateOption() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myDisk";
  const options = {
    body: {
      location: "West US",
      properties: {
        creationData: {
          createOption: "ImportSecure",
          securityDataUri: "https://mystorageaccount.blob.core.windows.net/osimages/vmgs.vhd",
          sourceUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
          storageAccountId:
            "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount",
        },
        osType: "Windows",
        securityProfile: {
          securityType: "ConfidentialVM_VMGuestStateOnlyEncryptedWithPlatformKey",
        },
      },
    },
    queryParameters: { "api-version": "2022-07-02" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}",
      subscriptionId,
      resourceGroupName,
      diskName,
    )
    .put(options);
  const poller = await getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createAManagedDiskFromImportSecureCreateOption().catch(console.error);
