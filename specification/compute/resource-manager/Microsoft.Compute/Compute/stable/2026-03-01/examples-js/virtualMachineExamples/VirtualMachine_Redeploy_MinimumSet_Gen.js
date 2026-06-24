const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to shuts down the virtual machine, moves it to a new node, and powers it back on.
 *
 * @summary shuts down the virtual machine, moves it to a new node, and powers it back on.
 * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Redeploy_MinimumSet_Gen.json
 */
async function virtualMachineRedeployMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.virtualMachines.redeploy("rgcompute", "aaaaaaaaaaaaaaa");
}
