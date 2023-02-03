const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a checkpoint in virtual machine.
 *
 * @summary Creates a checkpoint in virtual machine.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateCheckpointVirtualMachine.json
 */
async function createCheckpointVirtualMachine() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "testrg";
  const virtualMachineName = "DemoVM";
  const body = {
    name: "Demo Checkpoint name",
    description: "Demo Checkpoint description",
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.virtualMachines.beginCreateCheckpointAndWait(
    resourceGroupName,
    virtualMachineName,
    options
  );
  console.log(result);
}
