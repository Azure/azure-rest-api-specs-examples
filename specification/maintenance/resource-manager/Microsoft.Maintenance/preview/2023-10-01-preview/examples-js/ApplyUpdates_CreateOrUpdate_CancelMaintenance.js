const { MaintenanceManagementClient } = require("@azure/arm-maintenance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Apply maintenance updates to resource
 *
 * @summary Apply maintenance updates to resource
 * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/ApplyUpdates_CreateOrUpdate_CancelMaintenance.json
 */
async function applyUpdatesCreateOrUpdateOrCancel() {
  const subscriptionId =
    process.env["MAINTENANCE_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = process.env["MAINTENANCE_RESOURCE_GROUP"] || "examplerg";
  const providerName = "Microsoft.Maintenance";
  const resourceType = "maintenanceConfigurations";
  const resourceName = "maintenanceConfig1";
  const applyUpdateName = "20230901121200";
  const applyUpdate = { status: "Cancel" };
  const credential = new DefaultAzureCredential();
  const client = new MaintenanceManagementClient(credential, subscriptionId);
  const result = await client.applyUpdates.createOrUpdateOrCancel(
    resourceGroupName,
    providerName,
    resourceType,
    resourceName,
    applyUpdateName,
    applyUpdate,
  );
  console.log(result);
}
