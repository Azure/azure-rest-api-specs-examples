const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update a virtual hard disk.
 *
 * @summary The operation to update a virtual hard disk.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/UpdateVirtualHardDisk.json
 */
async function updateVirtualHardDisk() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const virtualHardDiskName = "test-vhd";
  const virtualHardDisks = {
    tags: { additionalProperties: "sample" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.virtualHardDisksOperations.beginUpdateAndWait(
    resourceGroupName,
    virtualHardDiskName,
    virtualHardDisks
  );
  console.log(result);
}
