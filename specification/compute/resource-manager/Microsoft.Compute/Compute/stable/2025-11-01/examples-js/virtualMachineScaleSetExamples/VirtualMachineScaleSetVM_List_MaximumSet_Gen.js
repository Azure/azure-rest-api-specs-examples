const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of all virtual machines in a VM scale sets.
 *
 * @summary gets a list of all virtual machines in a VM scale sets.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_List_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetVMListMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.virtualMachineScaleSetVMs.list(
    "rgcompute",
    "aaaaaaaaaaaaaaaaaaaaaa",
    { filter: "aaaaaaaaaaaaaa", select: "aaaaaaaaaaaaaaaaaaaaa", expand: "aaaaaaaaaaaaa" },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
