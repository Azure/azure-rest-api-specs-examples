const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a virtual machine instance. Please note some properties can be set only during virtual machine instance creation.
 *
 * @summary The operation to create or update a virtual machine instance. Please note some properties can be set only during virtual machine instance creation.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/PutVirtualMachineInstanceWithGalleryImage.json
 */
async function putVirtualMachineInstanceWithGalleryImage() {
  const resourceUri =
    "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM";
  const virtualMachineInstance = {
    extendedLocation: {
      name: "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location",
      type: "CustomLocation",
    },
    hardwareProfile: { vmSize: "Default" },
    networkProfile: { networkInterfaces: [{ id: "test-nic" }] },
    osProfile: {
      adminPassword: "password",
      adminUsername: "localadmin",
      computerName: "luamaster",
    },
    securityProfile: {
      enableTPM: true,
      uefiSettings: { secureBootEnabled: true },
    },
    storageProfile: {
      imageReference: {
        id: "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/galleryImages/test-gallery-image",
      },
      vmConfigStoragePathId:
        "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-container",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential);
  const result = await client.virtualMachineInstances.beginCreateOrUpdateAndWait(
    resourceUri,
    virtualMachineInstance
  );
  console.log(result);
}
