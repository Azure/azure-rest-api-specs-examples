const { ExtensionsClient } = require("@azure/arm-kubernetesconfiguration-extensions");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Patch an existing Kubernetes Cluster Extension.
 *
 * @summary Patch an existing Kubernetes Cluster Extension.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/extensions/stable/2024-11-01/examples/PatchExtension.json
 */
async function updateExtension() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subId1";
  const resourceGroupName = process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const extensionName = "ClusterMonitor";
  const patchExtension = {
    autoUpgradeMinorVersion: true,
    configurationProtectedSettings: { omsagentSecretKey: "secretKeyValue01" },
    configurationSettings: {
      omsagentEnvClusterName: "clusterName1",
      omsagentSecretWsid: "fakeTokenPlaceholder",
    },
    releaseTrain: "Preview",
  };
  const credential = new DefaultAzureCredential();
  const client = new ExtensionsClient(credential, subscriptionId);
  const result = await client.extensions.beginUpdateAndWait(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    extensionName,
    patchExtension,
  );
  console.log(result);
}
