const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a checkpoint in virtual machine.
 *
 * @summary Deletes a checkpoint in virtual machine.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteCheckpointVirtualMachine.json
 */
async function deleteCheckpointVirtualMachine() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "testrg";
  const virtualMachineName = "DemoVM";
  const body = { id: "Demo CheckpointID" };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.virtualMachines.beginDeleteCheckpointAndWait(
    resourceGroupName,
    virtualMachineName,
    options
  );
  console.log(result);
}

deleteCheckpointVirtualMachine().catch(console.error);
