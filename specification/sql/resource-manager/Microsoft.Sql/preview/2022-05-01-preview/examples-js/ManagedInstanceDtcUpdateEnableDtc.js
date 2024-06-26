const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates managed instance DTC settings.
 *
 * @summary Updates managed instance DTC settings.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/ManagedInstanceDtcUpdateEnableDtc.json
 */
async function updatesManagedInstanceDtcSettingsByEnablingDtc() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "testrg";
  const managedInstanceName = "testinstance";
  const dtcName = "current";
  const parameters = { dtcEnabled: true };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstanceDtcs.beginCreateOrUpdateAndWait(
    resourceGroupName,
    managedInstanceName,
    dtcName,
    parameters,
  );
  console.log(result);
}
