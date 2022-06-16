const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Execute a schedule. This operation can take a while to complete.
 *
 * @summary Execute a schedule. This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_Execute.json
 */
async function virtualMachineSchedulesExecute() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const virtualMachineName = "{vmName}";
  const name = "LabVmsShutdown";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.virtualMachineSchedules.beginExecuteAndWait(
    resourceGroupName,
    labName,
    virtualMachineName,
    name
  );
  console.log(result);
}

virtualMachineSchedulesExecute().catch(console.error);
