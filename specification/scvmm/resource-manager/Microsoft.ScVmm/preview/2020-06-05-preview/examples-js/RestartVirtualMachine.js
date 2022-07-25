const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restart virtual machine.
 *
 * @summary Restart virtual machine.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/RestartVirtualMachine.json
 */
async function restartVirtualMachine() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "testrg";
  const virtualMachineName = "DemoVM";
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.virtualMachines.beginRestartAndWait(
    resourceGroupName,
    virtualMachineName
  );
  console.log(result);
}

restartVirtualMachine().catch(console.error);
