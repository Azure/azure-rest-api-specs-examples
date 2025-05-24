const { ExtensionTypesClient } = require("@azure/arm-kubernetesconfiguration-extensiontypes");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List the version for an Extension Type installable to the cluster.
 *
 * @summary List the version for an Extension Type installable to the cluster.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/extensionTypes/preview/2024-11-01-preview/examples/ListExtensionTypeVersions.json
 */
async function listExtensionTypeVersions() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subId1";
  const resourceGroupName = process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "my-cluster";
  const extensionTypeName = "my-extension-type";
  const releaseTrain = "stable";
  const majorVersion = "2";
  const showLatest = true;
  const options = {
    releaseTrain,
    majorVersion,
    showLatest,
  };
  const credential = new DefaultAzureCredential();
  const client = new ExtensionTypesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.extensionTypes.listClusterListVersions(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    extensionTypeName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
