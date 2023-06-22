const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a SQL virtual machine.
 *
 * @summary Updates a SQL virtual machine.
 * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/UpdateSqlVirtualMachine.json
 */
async function updatesASqlVirtualMachineTags() {
  const subscriptionId =
    process.env["SQLVIRTUALMACHINE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQLVIRTUALMACHINE_RESOURCE_GROUP"] || "testrg";
  const sqlVirtualMachineName = "testvm";
  const parameters = { tags: { mytag: "myval" } };
  const credential = new DefaultAzureCredential();
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const result = await client.sqlVirtualMachines.beginUpdateAndWait(
    resourceGroupName,
    sqlVirtualMachineName,
    parameters
  );
  console.log(result);
}
