const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Captures the VM by copying virtual hard disks of the VM and outputs a template that can be used to create similar VMs.
 *
 * @summary Captures the VM by copying virtual hard disks of the VM and outputs a template that can be used to create similar VMs.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExamples/VirtualMachines_Capture_MinimumSet_Gen.json
 */
async function virtualMachinesCaptureMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmName = "aaaaaaaaaaaaa";
  const parameters = {
    destinationContainerName: "aaaaaaa",
    overwriteVhds: true,
    vhdPrefix: "aaaaaaaaa",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginCaptureAndWait(
    resourceGroupName,
    vmName,
    parameters
  );
  console.log(result);
}

virtualMachinesCaptureMinimumSetGen().catch(console.error);
