const { MaintenanceManagementClient } = require("@azure/arm-maintenance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Configuration records within a subscription and resource group
 *
 * @summary Get Configuration records within a subscription and resource group
 * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/MaintenanceConfigurationsResourceGroup_List.json
 */
async function maintenanceConfigurationsResourceGroupList() {
  const subscriptionId =
    process.env["MAINTENANCE_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = process.env["MAINTENANCE_RESOURCE_GROUP"] || "examplerg";
  const credential = new DefaultAzureCredential();
  const client = new MaintenanceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.maintenanceConfigurationsForResourceGroup.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
