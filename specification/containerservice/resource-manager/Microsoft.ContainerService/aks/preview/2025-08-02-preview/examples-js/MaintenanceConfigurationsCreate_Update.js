const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a maintenance configuration in the specified managed cluster.
 *
 * @summary Creates or updates a maintenance configuration in the specified managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-08-02-preview/examples/MaintenanceConfigurationsCreate_Update.json
 */
async function createOrUpdateMaintenanceConfiguration() {
  const subscriptionId =
    process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const resourceName = "clustername1";
  const configName = "default";
  const parameters = {
    notAllowedTime: [
      {
        end: new Date("2020-11-30T12:00:00Z"),
        start: new Date("2020-11-26T03:00:00Z"),
      },
    ],
    timeInWeek: [{ day: "Monday", hourSlots: [1, 2] }],
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.maintenanceConfigurations.createOrUpdate(
    resourceGroupName,
    resourceName,
    configName,
    parameters,
  );
  console.log(result);
}
