const { ExtensionsClient } = require("@azure/arm-kubernetesconfiguration-extensions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets Kubernetes Cluster Extension.
 *
 * @summary gets Kubernetes Cluster Extension.
 * x-ms-original-file: 2024-11-01/GetExtensionWithPlan.json
 */
async function getExtensionWithPlan() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId1";
  const client = new ExtensionsClient(credential, subscriptionId);
  const result = await client.extensions.get(
    "rg1",
    "Microsoft.Kubernetes",
    "connectedClusters",
    "clusterName1",
    "azureVote",
  );
  console.log(result);
}
