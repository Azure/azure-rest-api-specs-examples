const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of virtual machine image publishers for the specified Azure location.
 *
 * @summary Gets a list of virtual machine image publishers for the specified Azure location.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineImages_ListPublishers_MinimumSet_Gen.json
 */
async function virtualMachineImagesListPublishersMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.listPublishers(location);
  console.log(result);
}

virtualMachineImagesListPublishersMinimumSetGen().catch(console.error);
