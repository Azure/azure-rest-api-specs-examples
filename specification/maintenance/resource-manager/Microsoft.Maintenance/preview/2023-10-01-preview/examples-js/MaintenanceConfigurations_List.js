const { MaintenanceManagementClient } = require("@azure/arm-maintenance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Configuration records within a subscription
 *
 * @summary Get Configuration records within a subscription
 * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/MaintenanceConfigurations_List.json
 */
async function maintenanceConfigurationsList() {
  const subscriptionId =
    process.env["MAINTENANCE_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const credential = new DefaultAzureCredential();
  const client = new MaintenanceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.maintenanceConfigurations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
