const {
  FluxConfigurationClient,
} = require("@azure/arm-kubernetesconfiguration-fluxconfigurations");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get Async Operation status
 *
 * @summary Get Async Operation status
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/fluxConfigurations/stable/2025-04-01/examples/GetFluxConfigurationAsyncOperationStatus.json
 */
async function fluxConfigurationAsyncOperationStatusGet() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subId1";
  const resourceGroupName = process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const fluxConfigurationName = "srs-fluxconfig";
  const operationId = "99999999-9999-9999-9999-999999999999";
  const credential = new DefaultAzureCredential();
  const client = new FluxConfigurationClient(credential, subscriptionId);
  const result = await client.fluxConfigOperationStatus.get(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    fluxConfigurationName,
    operationId,
  );
  console.log(result);
}
