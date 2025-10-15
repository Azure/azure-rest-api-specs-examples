const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list DbSystemShape resources by SubscriptionLocationResource
 *
 * @summary list DbSystemShape resources by SubscriptionLocationResource
 * x-ms-original-file: 2025-09-01/DbSystemShapes_ListByLocation_MinimumSet_Gen.json
 */
async function listDbSystemShapesByLocationGeneratedByMinimumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.dbSystemShapes.listByLocation("eastus")) {
    resArray.push(item);
  }

  console.log(resArray);
}
