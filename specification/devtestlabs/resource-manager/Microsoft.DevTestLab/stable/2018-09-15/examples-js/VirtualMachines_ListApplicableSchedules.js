const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the applicable start/stop schedules, if any.
 *
 * @summary Lists the applicable start/stop schedules, if any.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_ListApplicableSchedules.json
 */
async function virtualMachinesListApplicableSchedules() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{vmName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.virtualMachines.listApplicableSchedules(
    resourceGroupName,
    labName,
    name
  );
  console.log(result);
}

virtualMachinesListApplicableSchedules().catch(console.error);
