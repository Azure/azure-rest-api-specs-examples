const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all availability group listeners in a SQL virtual machine group.
 *
 * @summary Lists all availability group listeners in a SQL virtual machine group.
 * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/ListByGroupAvailabilityGroupListener.json
 */
async function listsAllAvailabilityGroupListenersInASqlVirtualMachineGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const sqlVirtualMachineGroupName = "testvmgroup";
  const credential = new DefaultAzureCredential();
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availabilityGroupListeners.listByGroup(
    resourceGroupName,
    sqlVirtualMachineGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsAllAvailabilityGroupListenersInASqlVirtualMachineGroup().catch(console.error);
