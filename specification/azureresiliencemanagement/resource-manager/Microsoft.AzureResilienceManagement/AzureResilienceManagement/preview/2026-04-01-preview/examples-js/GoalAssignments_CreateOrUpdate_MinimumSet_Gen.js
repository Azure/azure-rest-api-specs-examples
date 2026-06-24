const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a GoalAssignment
 *
 * @summary create a GoalAssignment
 * x-ms-original-file: 2026-04-01-preview/GoalAssignments_CreateOrUpdate_MinimumSet_Gen.json
 */
async function goalAssignmentsCreateOrUpdateMinimumSet() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  await client.goalAssignments.createOrUpdate("sg1", "ga1", {
    properties: {
      goalTemplateId: "/providers/Microsoft.AzureResilienceManagement/goaltemplates/gt1",
      goalAssignmentType: "Resiliency",
    },
  });
}
