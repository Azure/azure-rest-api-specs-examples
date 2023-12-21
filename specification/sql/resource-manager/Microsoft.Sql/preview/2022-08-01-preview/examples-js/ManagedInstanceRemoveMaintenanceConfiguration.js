const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a managed instance.
 *
 * @summary Updates a managed instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedInstanceRemoveMaintenanceConfiguration.json
 */
async function removeMaintenancePolicyFromManagedInstanceSelectDefaultMaintenancePolicy() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "20D7082A-0FC7-4468-82BD-542694D5042B";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "testrg";
  const managedInstanceName = "testinstance";
  const parameters = {
    maintenanceConfigurationId:
      "/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_Default",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstances.beginUpdateAndWait(
    resourceGroupName,
    managedInstanceName,
    parameters,
  );
  console.log(result);
}
