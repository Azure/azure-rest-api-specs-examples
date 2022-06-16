const { SourceControlConfigurationClient } = require("@azure/arm-kubernetesconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This will delete the YAML file used to set up the Flux Configuration, thus stopping future sync from the source repo.
 *
 * @summary This will delete the YAML file used to set up the Flux Configuration, thus stopping future sync from the source repo.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/DeleteFluxConfiguration.json
 */
async function deleteFluxConfiguration() {
  const subscriptionId = "subId1";
  const resourceGroupName = "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const fluxConfigurationName = "srs-fluxconfig";
  const credential = new DefaultAzureCredential();
  const client = new SourceControlConfigurationClient(credential, subscriptionId);
  const result = await client.fluxConfigurations.beginDeleteAndWait(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    fluxConfigurationName
  );
  console.log(result);
}

deleteFluxConfiguration().catch(console.error);
