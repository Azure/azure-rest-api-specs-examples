const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a OracleSubscription
 *
 * @summary Create a OracleSubscription
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/oracleSubscriptions_create.json
 */
async function createOrUpdateOracleSubscription() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resource = {
    plan: {
      name: "plan1",
      product: "product1",
      promotionCode: "none",
      publisher: "publisher1",
      version: "alpha",
    },
    properties: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.oracleSubscriptions.beginCreateOrUpdateAndWait(resource);
  console.log(result);
}
