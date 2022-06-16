const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List schedules in a given virtual machine.
 *
 * @summary List schedules in a given virtual machine.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_List.json
 */
async function virtualMachineSchedulesList() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const virtualMachineName = "{vmName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineSchedules.list(
    resourceGroupName,
    labName,
    virtualMachineName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

virtualMachineSchedulesList().catch(console.error);
