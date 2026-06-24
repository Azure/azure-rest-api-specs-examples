const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of virtual machine image offers for the specified location and publisher.
 *
 * @summary gets a list of virtual machine image offers for the specified location and publisher.
 * x-ms-original-file: 2026-03-01/virtualMachineImageExamples/VirtualMachineImage_ListOffers_MinimumSet_Gen.json
 */
async function virtualMachineImageListOffersMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.listOffers(
    "aaaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaaaaaaa",
  );
  console.log(result);
}
