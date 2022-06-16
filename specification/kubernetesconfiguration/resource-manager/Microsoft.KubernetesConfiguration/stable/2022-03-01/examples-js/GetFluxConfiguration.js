const { SourceControlConfigurationClient } = require("@azure/arm-kubernetesconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details of the Flux Configuration.
 *
 * @summary Gets details of the Flux Configuration.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/GetFluxConfiguration.json
 */
async function getFluxConfiguration() {
  const subscriptionId = "subId1";
  const resourceGroupName = "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const fluxConfigurationName = "srs-fluxconfig";
  const credential = new DefaultAzureCredential();
  const client = new SourceControlConfigurationClient(credential, subscriptionId);
  const result = await client.fluxConfigurations.get(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    fluxConfigurationName
  );
  console.log(result);
}

getFluxConfiguration().catch(console.error);
