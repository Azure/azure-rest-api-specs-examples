const { MaintenanceManagementClient } = require("@azure/arm-maintenance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get configuration assignment for resource..
 *
 * @summary Get configuration assignment for resource..
 * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/ConfigurationAssignmentsForResourceGroup_Get.json
 */
async function configurationAssignmentsForResourceGroupGet() {
  const subscriptionId =
    process.env["MAINTENANCE_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = process.env["MAINTENANCE_RESOURCE_GROUP"] || "examplerg";
  const configurationAssignmentName = "workervmConfiguration";
  const credential = new DefaultAzureCredential();
  const client = new MaintenanceManagementClient(credential, subscriptionId);
  const result = await client.configurationAssignmentsForResourceGroup.get(
    resourceGroupName,
    configurationAssignmentName,
  );
  console.log(result);
}
