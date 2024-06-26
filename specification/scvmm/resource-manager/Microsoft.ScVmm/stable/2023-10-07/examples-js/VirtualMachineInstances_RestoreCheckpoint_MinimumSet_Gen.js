const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restores to a checkpoint in virtual machine instance.
 *
 * @summary Restores to a checkpoint in virtual machine instance.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineInstances_RestoreCheckpoint_MinimumSet_Gen.json
 */
async function virtualMachineInstancesRestoreCheckpointMinimumSet() {
  const resourceUri = "gtgclehcbsyave";
  const body = {};
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const result = await client.virtualMachineInstances.beginRestoreCheckpointAndWait(
    resourceUri,
    body,
  );
  console.log(result);
}
