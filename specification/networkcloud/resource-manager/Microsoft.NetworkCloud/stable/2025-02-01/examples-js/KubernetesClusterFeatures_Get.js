const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get properties of the provided the Kubernetes cluster feature.
 *
 * @summary Get properties of the provided the Kubernetes cluster feature.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/KubernetesClusterFeatures_Get.json
 */
async function getKubernetesClusterFeature() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const kubernetesClusterName = "kubernetesClusterName";
  const featureName = "featureName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.kubernetesClusterFeatures.get(
    resourceGroupName,
    kubernetesClusterName,
    featureName,
  );
  console.log(result);
}
