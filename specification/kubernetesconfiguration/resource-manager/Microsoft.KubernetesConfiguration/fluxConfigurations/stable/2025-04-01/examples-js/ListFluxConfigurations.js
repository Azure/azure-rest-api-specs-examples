const {
  FluxConfigurationClient,
} = require("@azure/arm-kubernetesconfiguration-fluxconfigurations");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List all Flux Configurations.
 *
 * @summary List all Flux Configurations.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/fluxConfigurations/stable/2025-04-01/examples/ListFluxConfigurations.json
 */
async function listFluxConfiguration() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subId1";
  const resourceGroupName = process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const credential = new DefaultAzureCredential();
  const client = new FluxConfigurationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.fluxConfigurations.list(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
