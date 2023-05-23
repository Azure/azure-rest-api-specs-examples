const { SourceControlConfigurationClient } = require("@azure/arm-kubernetesconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Kubernetes Flux Configuration.
 *
 * @summary Create a new Kubernetes Flux Configuration.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/CreateFluxConfigurationWithBucket.json
 */
async function createFluxConfigurationWithBucketSourceKind() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subId1";
  const resourceGroupName = process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const fluxConfigurationName = "srs-fluxconfig";
  const fluxConfiguration = {
    bucket: {
      accessKey: "fluxminiotest",
      bucketName: "flux",
      syncIntervalInSeconds: 1000,
      timeoutInSeconds: 1000,
      url: "https://fluxminiotest.az.minio.io",
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
    sourceKind: "Bucket",
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
