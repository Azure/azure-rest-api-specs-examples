const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a maintenance configuration in the specified managed cluster.
 *
 * @summary creates or updates a maintenance configuration in the specified managed cluster.
 * x-ms-original-file: 2026-06-02-preview/MaintenanceConfigurationsCreate_LinkedMaintenanceWindow.json
 */
async function createALinkedMaintenanceConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.maintenanceConfigurations.createOrUpdate(
    "rg1",
    "clustername1",
    "aksManagedAutoUpgradeSchedule",
    {
      maintenanceWindowId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/maintenanceWindows/myMaintenanceWindow",
    },
  );
  console.log(result);
}
