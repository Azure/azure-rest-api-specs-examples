const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to perform add storage capacity on exadata infra
 *
 * @summary perform add storage capacity on exadata infra
 * x-ms-original-file: 2025-09-01/CloudExadataInfrastructures_AddStorageCapacity_MaximumSet_Gen.json
 */
async function performAddStorageCapacityOnExadataInfraGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.cloudExadataInfrastructures.addStorageCapacity(
    "rgopenapi",
    "Replace this value with a string matching RegExp .*",
  );
  console.log(result);
}
