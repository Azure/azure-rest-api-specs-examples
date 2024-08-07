const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a CloudExadataInfrastructure
 *
 * @summary Update a CloudExadataInfrastructure
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/exaInfra_patch.json
 */
async function patchExadataInfrastructure() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const cloudexadatainfrastructurename = "infra1";
  const properties = {};
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.cloudExadataInfrastructures.beginUpdateAndWait(
    resourceGroupName,
    cloudexadatainfrastructurename,
    properties,
  );
  console.log(result);
}
