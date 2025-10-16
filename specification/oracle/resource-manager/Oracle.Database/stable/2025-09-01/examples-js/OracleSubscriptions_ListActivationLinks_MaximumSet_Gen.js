const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list Activation Links
 *
 * @summary list Activation Links
 * x-ms-original-file: 2025-09-01/OracleSubscriptions_ListActivationLinks_MaximumSet_Gen.json
 */
async function listActivationLinksForTheOracleSubscriptionGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.oracleSubscriptions.listActivationLinks();
  console.log(result);
}
