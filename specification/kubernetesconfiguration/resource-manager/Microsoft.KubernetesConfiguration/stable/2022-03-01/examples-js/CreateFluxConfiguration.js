const { SourceControlConfigurationClient } = require("@azure/arm-kubernetesconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Kubernetes Flux Configuration.
 *
 * @summary Create a new Kubernetes Flux Configuration.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/CreateFluxConfiguration.json
 */
async function createFluxConfiguration() {
  const subscriptionId = "subId1";
  const resourceGroupName = "rg1";
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
    scope: "cluster",
    sourceKind: "GitRepository",
    suspend: false,
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

createFluxConfiguration().catch(console.error);
