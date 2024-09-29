const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified maintenance configuration of a managed cluster.
 *
 * @summary Gets the specified maintenance configuration of a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-07-02-preview/examples/MaintenanceConfigurationsGet.json
 */
async function getMaintenanceConfiguration() {
  const subscriptionId =
    process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const resourceName = "clustername1";
  const configName = "default";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.maintenanceConfigurations.get(
    resourceGroupName,
    resourceName,
    configName,
  );
  console.log(result);
}
