const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a GiMinorVersion
 *
 * @summary get a GiMinorVersion
 * x-ms-original-file: 2025-03-01/GiMinorVersions_Get_MaximumSet_Gen.json
 */
async function giMinorVersionsGetMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.giMinorVersions.get("eastus", "giVersionName", "giMinorVersionName");
  console.log(result);
}
