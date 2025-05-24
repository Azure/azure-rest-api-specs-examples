const {
  FluxConfigurationClient,
} = require("@azure/arm-kubernetesconfiguration-fluxconfigurations");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets details of the Flux Configuration.
 *
 * @summary Gets details of the Flux Configuration.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/fluxConfigurations/stable/2025-04-01/examples/GetFluxConfiguration.json
 */
async function getFluxConfiguration() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subId1";
  const resourceGroupName = process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const fluxConfigurationName = "srs-fluxconfig";
  const credential = new DefaultAzureCredential();
  const client = new FluxConfigurationClient(credential, subscriptionId);
  const result = await client.fluxConfigurations.get(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    fluxConfigurationName,
  );
  console.log(result);
}
