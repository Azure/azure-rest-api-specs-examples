const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified maintenance configuration of a managed cluster.
 *
 * @summary Gets the specified maintenance configuration of a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-07-01/examples/MaintenanceConfigurationsGet.json
 */
async function getMaintenanceConfiguration() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const configName = "default";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.maintenanceConfigurations.get(
    resourceGroupName,
    resourceName,
    configName
  );
  console.log(result);
}

getMaintenanceConfiguration().catch(console.error);
