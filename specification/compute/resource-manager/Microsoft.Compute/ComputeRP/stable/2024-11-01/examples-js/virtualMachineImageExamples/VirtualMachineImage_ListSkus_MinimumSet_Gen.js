const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets a list of virtual machine image SKUs for the specified location, publisher, and offer.
 *
 * @summary Gets a list of virtual machine image SKUs for the specified location, publisher, and offer.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/virtualMachineImageExamples/VirtualMachineImage_ListSkus_MinimumSet_Gen.json
 */
async function virtualMachineImageListSkusMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const location = "aaaa";
  const publisherName = "aaaaaaaaaaaaa";
  const offer = "aaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.listSkus(location, publisherName, offer);
  console.log(result);
}
