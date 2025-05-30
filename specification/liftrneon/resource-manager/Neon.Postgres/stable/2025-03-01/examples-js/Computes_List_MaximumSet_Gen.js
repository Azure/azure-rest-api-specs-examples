const { PostgresClient } = require("@azure/arm-neonpostgres");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list Compute resources by Branch
 *
 * @summary list Compute resources by Branch
 * x-ms-original-file: 2025-03-01/Computes_List_MaximumSet_Gen.json
 */
async function computesListMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "9B8E3300-C5FA-442B-A259-3F6F614D5BD4";
  const client = new PostgresClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.computes.list(
    "rgneon",
    "test-org",
    "entity-name",
    "entity-name",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
