const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of virtual machine image publishers for the specified Azure location.
 *
 * @summary Gets a list of virtual machine image publishers for the specified Azure location.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImage_ListPublishers_MaximumSet_Gen.json
 */
async function virtualMachineImageListPublishersMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const location = "aaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.listPublishers(location);
  console.log(result);
}
