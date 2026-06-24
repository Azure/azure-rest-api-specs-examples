const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of virtual machine image publishers for the specified Azure location.
 *
 * @summary gets a list of virtual machine image publishers for the specified Azure location.
 * x-ms-original-file: 2026-03-01/virtualMachineImageExamples/VirtualMachineImage_ListPublishers_MaximumSet_Gen.json
 */
async function virtualMachineImageListPublishersMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.listPublishers("aaaaa");
  console.log(result);
}
