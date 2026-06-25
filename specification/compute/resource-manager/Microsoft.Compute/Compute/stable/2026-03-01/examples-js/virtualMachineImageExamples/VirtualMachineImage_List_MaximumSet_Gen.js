const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
 *
 * @summary gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
 * x-ms-original-file: 2026-03-01/virtualMachineImageExamples/VirtualMachineImage_List_MaximumSet_Gen.json
 */
async function virtualMachineImageListMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.list(
    "aaaaaaaaaaaaaaa",
    "aaaaaa",
    "aaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaa",
    { expand: "aaaaaaaaaaaaaaaaaaaaaaaa", top: 18, orderby: "aa" },
  );
  console.log(result);
}
