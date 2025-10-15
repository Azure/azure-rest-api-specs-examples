const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a ExascaleDbStorageVault
 *
 * @summary get a ExascaleDbStorageVault
 * x-ms-original-file: 2025-09-01/ExascaleDbStorageVaults_Get_MinimumSet_Gen.json
 */
async function exascaleDbStorageVaultsGetMaximumSetGeneratedByMinimumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.exascaleDbStorageVaults.get("rgopenapi", "storagevault1");
  console.log(result);
}
