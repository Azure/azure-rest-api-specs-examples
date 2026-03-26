const {
  FluxConfigurationClient,
} = require("@azure/arm-kubernetesconfiguration-fluxconfigurations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a new Kubernetes Flux Configuration.
 *
 * @summary create a new Kubernetes Flux Configuration.
 * x-ms-original-file: 2025-04-01/CreateFluxConfigurationWithBucket.json
 */
async function createFluxConfigurationWithBucketSourceKind() {
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
      bucket: {
        accessKey: "fluxminiotest",
        bucketName: "flux",
        syncIntervalInSeconds: 1000,
        timeoutInSeconds: 1000,
        url: "https://fluxminiotest.az.minio.io",
      },
      kustomizations: {
        "srs-kustomization1": {
          path: "./test/path",
          dependsOn: [],
          syncIntervalInSeconds: 600,
          timeoutInSeconds: 600,
        },
        "srs-kustomization2": {
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
      sourceKind: "Bucket",
      suspend: false,
    },
  );
  console.log(result);
}
