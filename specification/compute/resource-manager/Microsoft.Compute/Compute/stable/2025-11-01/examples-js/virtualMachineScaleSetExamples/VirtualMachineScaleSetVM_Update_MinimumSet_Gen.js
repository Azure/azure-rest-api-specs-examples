const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a virtual machine of a VM scale set.
 *
 * @summary updates a virtual machine of a VM scale set.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Update_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetVMUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.update(
    "rgcompute",
    "aaaaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaa",
    { location: "westus" },
  );
  console.log(result);
}
