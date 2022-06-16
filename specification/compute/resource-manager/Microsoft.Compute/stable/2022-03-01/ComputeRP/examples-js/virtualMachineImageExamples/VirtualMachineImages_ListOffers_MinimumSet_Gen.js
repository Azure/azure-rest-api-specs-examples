const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of virtual machine image offers for the specified location and publisher.
 *
 * @summary Gets a list of virtual machine image offers for the specified location and publisher.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineImageExamples/VirtualMachineImages_ListOffers_MinimumSet_Gen.json
 */
async function virtualMachineImagesListOffersMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaaaaaaaaaaaa";
  const publisherName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.listOffers(location, publisherName);
  console.log(result);
}

virtualMachineImagesListOffersMinimumSetGen().catch(console.error);
