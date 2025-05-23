const { ExtensionTypesClient } = require("@azure/arm-kubernetesconfiguration-extensiontypes");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get details of a version for an Extension Type installable to the cluster.
 *
 * @summary Get details of a version for an Extension Type installable to the cluster.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/extensionTypes/preview/2024-11-01-preview/examples/GetExtensionTypeVersion.json
 */
async function listExtensionTypeVersions() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subId1";
  const resourceGroupName = process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "my-cluster";
  const extensionTypeName = "my-extension-type";
  const versionNumber = "v1.3.2";
  const credential = new DefaultAzureCredential();
  const client = new ExtensionTypesClient(credential, subscriptionId);
  const result = await client.extensionTypes.clusterGetVersion(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    extensionTypeName,
    versionNumber,
  );
  console.log(result);
}
