const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a ExascaleDbStorageVault
 *
 * @summary update a ExascaleDbStorageVault
 * x-ms-original-file: 2025-09-01/ExascaleDbStorageVaults_Update_MaximumSet_Gen.json
 */
async function exascaleDbStorageVaultsUpdateMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.exascaleDbStorageVaults.update("rgopenapi", "storagevault1", {
    tags: { key6486: "fxbuboilsiapodppdtfls" },
  });
  console.log(result);
}
