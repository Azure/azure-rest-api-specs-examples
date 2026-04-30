const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
 *
 * @summary gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
 * x-ms-original-file: 2025-11-01/virtualMachineImageExamples/VirtualMachineImage_List_MinimumSet_Gen.json
 */
async function virtualMachineImageListMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.list(
    "aaaaaaa",
    "aaaaaaaaaaa",
    "aaaaaaaaaa",
    "aaaaaa",
  );
  console.log(result);
}
