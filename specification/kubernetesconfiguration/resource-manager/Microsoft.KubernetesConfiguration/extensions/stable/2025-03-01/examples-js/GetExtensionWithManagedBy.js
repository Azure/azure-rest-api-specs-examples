const { ExtensionsClient } = require("@azure/arm-kubernetesconfiguration-extensions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets Kubernetes Cluster Extension.
 *
 * @summary gets Kubernetes Cluster Extension.
 * x-ms-original-file: 2025-03-01/GetExtensionWithManagedBy.json
 */
async function getExtensionWithManagedBy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId1";
  const client = new ExtensionsClient(credential, subscriptionId);
  const result = await client.extensions.get(
    "rg1",
    "Microsoft.ContainerService",
    "managedClusters",
    "clusterName1",
    "azureVote",
  );
  console.log(result);
}
