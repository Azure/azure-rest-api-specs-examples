const { SourceControlConfigurationClient } = require("@azure/arm-kubernetesconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Kubernetes Flux Configuration.
 *
 * @summary Create a new Kubernetes Flux Configuration.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/CreateFluxConfiguration.json
 */
async function createFluxConfiguration() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subId1";
  const resourceGroupName = process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const fluxConfigurationName = "srs-fluxconfig";
  const fluxConfiguration = {
    gitRepository: {
      httpsCACert: "ZXhhbXBsZWNlcnRpZmljYXRl",
      repositoryRef: { branch: "master" },
      syncIntervalInSeconds: 600,
      timeoutInSeconds: 600,
      url: "https://github.com/Azure/arc-k8s-demo",
    },
    kustomizations: {
      srsKustomization1: {
        path: "./test/path",
        dependsOn: [],
        postBuild: {
          substitute: { clusterEnv: "prod", replicaCount: "2" },
          substituteFrom: [{ name: "cluster-test", kind: "ConfigMap", optional: true }],
        },
        syncIntervalInSeconds: 600,
        timeoutInSeconds: 600,
        wait: true,
      },
      srsKustomization2: {
        path: "./other/test/path",
        dependsOn: ["srs-kustomization1"],
        postBuild: {
          substituteFrom: [
            { name: "cluster-values", kind: "ConfigMap", optional: true },
            { name: "secret-name", kind: "Secret", optional: false },
          ],
        },
        prune: false,
        retryIntervalInSeconds: 600,
        syncIntervalInSeconds: 600,
        timeoutInSeconds: 600,
        wait: false,
      },
    },
    namespace: "srs-namespace",
    reconciliationWaitDuration: "PT30M",
    scope: "cluster",
    sourceKind: "GitRepository",
    suspend: false,
    waitForReconciliation: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new SourceControlConfigurationClient(credential, subscriptionId);
  const result = await client.fluxConfigurations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    fluxConfigurationName,
    fluxConfiguration
  );
  console.log(result);
}
