const { MaintenanceManagementClient } = require("@azure/arm-maintenance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Public Maintenance Configuration record
 *
 * @summary Get Public Maintenance Configuration record
 * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/PublicMaintenanceConfigurations_GetForResource.json
 */
async function publicMaintenanceConfigurationsGetForResource() {
  const subscriptionId =
    process.env["MAINTENANCE_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceName = "configuration1";
  const credential = new DefaultAzureCredential();
  const client = new MaintenanceManagementClient(credential, subscriptionId);
  const result = await client.publicMaintenanceConfigurations.get(resourceName);
  console.log(result);
}
