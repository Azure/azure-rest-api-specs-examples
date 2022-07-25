const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all SQL virtual machine groups in a subscription.
 *
 * @summary Gets all SQL virtual machine groups in a subscription.
 * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/ListSubscriptionSqlVirtualMachineGroup.json
 */
async function getsAllSqlVirtualMachineGroupsInASubscription() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const credential = new DefaultAzureCredential();
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sqlVirtualMachineGroups.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsAllSqlVirtualMachineGroupsInASubscription().catch(console.error);
