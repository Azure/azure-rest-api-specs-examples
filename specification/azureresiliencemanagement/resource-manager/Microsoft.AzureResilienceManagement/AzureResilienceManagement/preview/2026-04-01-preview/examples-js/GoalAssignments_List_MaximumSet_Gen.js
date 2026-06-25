const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list GoalAssignment resources by tenant
 *
 * @summary list GoalAssignment resources by tenant
 * x-ms-original-file: 2026-04-01-preview/GoalAssignments_List_MaximumSet_Gen.json
 */
async function goalAssignmentsListMaximumSet() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.goalAssignments.list("zldmpkvqzifygkqau", {
    skipToken: "xntbyoswztnmvitj",
    top: 69,
  })) {
    resArray.push(item);
  }

  console.log(resArray);
}
