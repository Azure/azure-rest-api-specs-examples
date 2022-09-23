const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
 *
 * @summary Gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineImageExamples/VirtualMachineImages_List_MinimumSet_Gen.json
 */
async function virtualMachineImagesListMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaa";
  const publisherName = "aaaaaaaaaaa";
  const offer = "aaaaaaaaaa";
  const skus = "aaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.list(location, publisherName, offer, skus);
  console.log(result);
}

virtualMachineImagesListMinimumSetGen().catch(console.error);
