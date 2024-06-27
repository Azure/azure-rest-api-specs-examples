const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a checkpoint in virtual machine instance.
 *
 * @summary Deletes a checkpoint in virtual machine instance.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineInstances_DeleteCheckpoint_MaximumSet_Gen.json
 */
async function virtualMachineInstancesDeleteCheckpointMaximumSet() {
  const resourceUri = "gtgclehcbsyave";
  const body = {
    id: "eenfflimcbgqfsebdusophahjpk",
  };
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const result = await client.virtualMachineInstances.beginDeleteCheckpointAndWait(
    resourceUri,
    body,
  );
  console.log(result);
}
