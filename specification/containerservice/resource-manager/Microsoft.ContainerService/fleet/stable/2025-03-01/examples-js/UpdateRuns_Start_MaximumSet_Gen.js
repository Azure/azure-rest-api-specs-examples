const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to starts an UpdateRun.
 *
 * @summary starts an UpdateRun.
 * x-ms-original-file: 2025-03-01/UpdateRuns_Start_MaximumSet_Gen.json
 */
async function startsAnUpdateRunGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.updateRuns.start("rgfleets", "fleet1", "fleet1", {
    ifMatch: "bvhjlqeindkmljbbiypbqiaqgtkhlu",
  });
  console.log(result);
}
