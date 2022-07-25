const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Uninstalls and reinstalls the SQL Iaas Extension.
 *
 * @summary Uninstalls and reinstalls the SQL Iaas Extension.
 * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/RedeploySqlVirtualMachine.json
 */
async function uninstallsAndReinstallsTheSqlIaasExtension() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const sqlVirtualMachineName = "testvm";
  const credential = new DefaultAzureCredential();
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const result = await client.sqlVirtualMachines.beginRedeployAndWait(
    resourceGroupName,
    sqlVirtualMachineName
  );
  console.log(result);
}

uninstallsAndReinstallsTheSqlIaasExtension().catch(console.error);
