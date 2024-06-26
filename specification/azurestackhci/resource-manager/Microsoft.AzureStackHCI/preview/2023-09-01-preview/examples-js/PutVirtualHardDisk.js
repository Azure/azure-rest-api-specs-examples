const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a virtual hard disk. Please note some properties can be set only during virtual hard disk creation.
 *
 * @summary The operation to create or update a virtual hard disk. Please note some properties can be set only during virtual hard disk creation.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/PutVirtualHardDisk.json
 */
async function putVirtualHardDisk() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const virtualHardDiskName = "test-vhd";
  const virtualHardDisks = {
    diskSizeGB: 32,
    extendedLocation: {
      name: "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location",
      type: "CustomLocation",
    },
    location: "West US2",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.virtualHardDisksOperations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualHardDiskName,
    virtualHardDisks
  );
  console.log(result);
}
