const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List AutonomousDbVersion resources by Location
 *
 * @summary List AutonomousDbVersion resources by Location
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabaseVersion_listByLocation.json
 */
async function listAnAutonomousVersionsByLocation() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.autonomousDatabaseVersions.listByLocation(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
