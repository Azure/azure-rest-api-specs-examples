const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Saas Subscription Details
 *
 * @summary List Saas Subscription Details
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/oracleSubscriptions_listSaasSubscriptionDetails.json
 */
async function listSaasSubscriptionDetailsForTheOracleSubscription() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.oracleSubscriptions.beginListSaasSubscriptionDetailsAndWait();
  console.log(result);
}
