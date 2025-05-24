const {
  FluxConfigurationClient,
} = require("@azure/arm-kubernetesconfiguration-fluxconfigurations");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create a new Kubernetes Flux Configuration.
 *
 * @summary Create a new Kubernetes Flux Configuration.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/fluxConfigurations/stable/2025-04-01/examples/CreateFluxConfigurationWithOCIRepository.json
 */
async function createFluxConfigurationWithOciRepositorySourceKind() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subId1";
  const resourceGroupName = process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const fluxConfigurationName = "srs-fluxconfig";
  const fluxConfiguration = {
    kustomizations: {
      srsKustomization1: {
        path: "./test/path",
        dependsOn: [],
        syncIntervalInSeconds: 600,
        timeoutInSeconds: 600,
      },
      srsKustomization2: {
        path: "./other/test/path",
        dependsOn: ["srs-kustomization1"],
        prune: false,
        retryIntervalInSeconds: 600,
        syncIntervalInSeconds: 600,
        timeoutInSeconds: 600,
      },
    },
    namespace: "srs-namespace",
    ociRepository: {
      serviceAccountName: "testserviceaccount",
      syncIntervalInSeconds: 1000,
      timeoutInSeconds: 1000,
      url: "oci://ghcr.io/stefanprodan/manifests/podinfo",
    },
    scope: "cluster",
    sourceKind: "OCIRepository",
    suspend: false,
  };
  const credential = new DefaultAzureCredential();
  const client = new FluxConfigurationClient(credential, subscriptionId);
  const result = await client.fluxConfigurations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    fluxConfigurationName,
    fluxConfiguration,
  );
  console.log(result);
}
