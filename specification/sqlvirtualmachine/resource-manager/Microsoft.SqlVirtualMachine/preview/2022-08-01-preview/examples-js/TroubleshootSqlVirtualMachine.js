const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts SQL virtual machine troubleshooting.
 *
 * @summary Starts SQL virtual machine troubleshooting.
 * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/TroubleshootSqlVirtualMachine.json
 */
async function startSqlVirtualMachineTroubleshootingOperation() {
  const subscriptionId =
    process.env["SQLVIRTUALMACHINE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQLVIRTUALMACHINE_RESOURCE_GROUP"] || "testrg";
  const sqlVirtualMachineName = "testvm";
  const parameters = {
    endTimeUtc: new Date("2022-07-09T22:10:00Z"),
    properties: { unhealthyReplicaInfo: { availabilityGroupName: "AG1" } },
    startTimeUtc: new Date("2022-07-09T17:10:00Z"),
    troubleshootingScenario: "UnhealthyReplica",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const result = await client.sqlVirtualMachineTroubleshoot.beginTroubleshootAndWait(
    resourceGroupName,
    sqlVirtualMachineName,
    parameters
  );
  console.log(result);
}
