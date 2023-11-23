const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of maintenance configurations in the specified managed cluster.
 *
 * @summary Gets a list of maintenance configurations in the specified managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-09-02-preview/examples/MaintenanceConfigurationsList.json
 */
async function listMaintenanceConfigurationsByManagedCluster() {
  const subscriptionId =
    process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.maintenanceConfigurations.listByManagedCluster(
    resourceGroupName,
    resourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
