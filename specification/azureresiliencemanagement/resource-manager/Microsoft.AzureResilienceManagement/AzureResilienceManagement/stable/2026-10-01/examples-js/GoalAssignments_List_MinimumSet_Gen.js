const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists goal assignments in a service group.
 *
 * @summary lists goal assignments in a service group.
 * x-ms-original-file: 2026-10-01/GoalAssignments_List_MinimumSet_Gen.json
 */
async function goalAssignmentsListMinimumSet() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.goalAssignments.list("production-sg")) {
    resArray.push(item);
  }

  console.log(resArray);
}
