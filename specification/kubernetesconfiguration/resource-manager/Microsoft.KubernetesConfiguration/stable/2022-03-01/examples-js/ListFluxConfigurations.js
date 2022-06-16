const { SourceControlConfigurationClient } = require("@azure/arm-kubernetesconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all Flux Configurations.
 *
 * @summary List all Flux Configurations.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/ListFluxConfigurations.json
 */
async function listFluxConfiguration() {
  const subscriptionId = "subId1";
  const resourceGroupName = "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const credential = new DefaultAzureCredential();
  const client = new SourceControlConfigurationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fluxConfigurations.list(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listFluxConfiguration().catch(console.error);
