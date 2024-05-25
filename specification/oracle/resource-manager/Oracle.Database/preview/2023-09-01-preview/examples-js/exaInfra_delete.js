const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a CloudExadataInfrastructure
 *
 * @summary Delete a CloudExadataInfrastructure
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/exaInfra_delete.json
 */
async function deleteExadataInfrastructure() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const cloudexadatainfrastructurename = "infra1";
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.cloudExadataInfrastructures.beginDeleteAndWait(
    resourceGroupName,
    cloudexadatainfrastructurename,
  );
  console.log(result);
}
