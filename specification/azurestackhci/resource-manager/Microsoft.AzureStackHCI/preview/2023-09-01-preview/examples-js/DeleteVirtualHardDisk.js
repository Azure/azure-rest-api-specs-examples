const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete a virtual hard disk.
 *
 * @summary The operation to delete a virtual hard disk.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/DeleteVirtualHardDisk.json
 */
async function deleteVirtualHardDisk() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const virtualHardDiskName = "test-vhd";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.virtualHardDisksOperations.beginDeleteAndWait(
    resourceGroupName,
    virtualHardDiskName
  );
  console.log(result);
}
