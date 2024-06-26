const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete a virtual machine instance.
 *
 * @summary The operation to delete a virtual machine instance.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineInstances_Delete_MaximumSet_Gen.json
 */
async function virtualMachineInstancesDeleteMaximumSet() {
  const resourceUri = "gtgclehcbsyave";
  const force = "true";
  const deleteFromHost = "true";
  const options = {
    force,
    deleteFromHost,
  };
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const result = await client.virtualMachineInstances.beginDeleteAndWait(resourceUri, options);
  console.log(result);
}
