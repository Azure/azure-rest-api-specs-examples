const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified maintenance configuration of a managed cluster.
 *
 * @summary gets the specified maintenance configuration of a managed cluster.
 * x-ms-original-file: 2026-06-02-preview/MaintenanceConfigurationsGet_LinkedMaintenanceWindow.json
 */
async function getALinkedMaintenanceConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.maintenanceConfigurations.get(
    "rg1",
    "clustername1",
    "aksManagedAutoUpgradeSchedule",
  );
  console.log(result);
}
