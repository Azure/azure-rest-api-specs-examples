const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list GoalTemplate resources by tenant
 *
 * @summary list GoalTemplate resources by tenant
 * x-ms-original-file: 2026-04-01-preview/GoalTemplates_List_MinimumSet_Gen.json
 */
async function goalTemplatesListMinimumSet() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.goalTemplates.list("sg1")) {
    resArray.push(item);
  }

  console.log(resArray);
}
