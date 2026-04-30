const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to captures the VM by copying virtual hard disks of the VM and outputs a template that can be used to create similar VMs.
 *
 * @summary captures the VM by copying virtual hard disks of the VM and outputs a template that can be used to create similar VMs.
 * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Capture_MaximumSet_Gen.json
 */
async function virtualMachineCaptureMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.capture("rgcompute", "aaaaaaaaaaaaaaaaaaaa", {
    vhdPrefix: "aaaaaaaaa",
    destinationContainerName: "aaaaaaa",
    overwriteVhds: true,
  });
  console.log(result);
}
