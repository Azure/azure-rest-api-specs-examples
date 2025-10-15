const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a OracleSubscription
 *
 * @summary update a OracleSubscription
 * x-ms-original-file: 2025-09-01/OracleSubscriptions_Update_MinimumSet_Gen.json
 */
async function patchOracleSubscriptionGeneratedByMinimumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.oracleSubscriptions.update({});
  console.log(result);
}
