const {
  FluxConfigurationClient,
} = require("@azure/arm-kubernetesconfiguration-fluxconfigurations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a new Kubernetes Flux Configuration.
 *
 * @summary create a new Kubernetes Flux Configuration.
 * x-ms-original-file: 2025-04-01/CreateFluxConfiguration.json
 */
async function createFluxConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId1";
  const client = new FluxConfigurationClient(credential, subscriptionId);
  const result = await client.fluxConfigurations.createOrUpdate(
    "rg1",
    "Microsoft.Kubernetes",
    "connectedClusters",
    "clusterName1",
    "srs-fluxconfig",
    {
      gitRepository: {
        httpsCACert: "ZXhhbXBsZWNlcnRpZmljYXRl",
        repositoryRef: { branch: "master" },
        syncIntervalInSeconds: 600,
        timeoutInSeconds: 600,
        url: "https://github.com/Azure/arc-k8s-demo",
      },
      kustomizations: {
        "srs-kustomization1": {
          path: "./test/path",
          dependsOn: [],
          postBuild: {
            substitute: { cluster_env: "prod", replica_count: "2" },
            substituteFrom: [{ name: "cluster-test", kind: "ConfigMap", optional: true }],
          },
          syncIntervalInSeconds: 600,
          timeoutInSeconds: 600,
          wait: true,
        },
        "srs-kustomization2": {
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
    },
  );
  console.log(result);
}
