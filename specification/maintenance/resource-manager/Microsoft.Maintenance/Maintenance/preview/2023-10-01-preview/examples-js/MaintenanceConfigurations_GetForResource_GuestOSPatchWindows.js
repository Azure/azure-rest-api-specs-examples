const { MaintenanceManagementClient } = require("@azure/arm-maintenance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get Configuration record
 *
 * @summary get Configuration record
 * x-ms-original-file: 2023-10-01-preview/MaintenanceConfigurations_GetForResource_GuestOSPatchWindows.json
 */
async function maintenanceConfigurationsGetForResourceGuestOSPatchWindows() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const client = new MaintenanceManagementClient(credential, subscriptionId);
  const result = await client.maintenanceConfigurations.get("examplerg", "configuration1");
  console.log(result);
}
