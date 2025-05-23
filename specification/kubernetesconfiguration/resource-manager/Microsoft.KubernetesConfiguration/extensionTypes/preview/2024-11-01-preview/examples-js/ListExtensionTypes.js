const { ExtensionTypesClient } = require("@azure/arm-kubernetesconfiguration-extensiontypes");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List installable Extension Types for the cluster based region and type for the cluster.
 *
 * @summary List installable Extension Types for the cluster based region and type for the cluster.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/extensionTypes/preview/2024-11-01-preview/examples/ListExtensionTypes.json
 */
async function getExtensionTypes() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subId1";
  const resourceGroupName = process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "my-cluster";
  const publisherId = "myPublisherId";
  const offerId = "myOfferId";
  const planId = "myPlanId";
  const releaseTrain = "stable";
  const options = {
    publisherId,
    offerId,
    planId,
    releaseTrain,
  };
  const credential = new DefaultAzureCredential();
  const client = new ExtensionTypesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.extensionTypes.list(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
