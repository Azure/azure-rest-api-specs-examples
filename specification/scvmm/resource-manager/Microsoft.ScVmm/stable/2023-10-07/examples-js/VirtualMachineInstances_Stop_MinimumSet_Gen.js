const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to power off (stop) a virtual machine instance.
 *
 * @summary The operation to power off (stop) a virtual machine instance.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineInstances_Stop_MinimumSet_Gen.json
 */
async function virtualMachineInstancesStopMinimumSet() {
  const resourceUri = "gtgclehcbsyave";
  const body = {};
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const result = await client.virtualMachineInstances.beginStopAndWait(resourceUri, body);
  console.log(result);
}
