const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update a virtual machine instance.
 *
 * @summary The operation to update a virtual machine instance.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineInstances_Update_MinimumSet_Gen.json
 */
async function virtualMachineInstancesUpdateMinimumSet() {
  const resourceUri = "gtgclehcbsyave";
  const properties = {};
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const result = await client.virtualMachineInstances.beginUpdateAndWait(resourceUri, properties);
  console.log(result);
}
