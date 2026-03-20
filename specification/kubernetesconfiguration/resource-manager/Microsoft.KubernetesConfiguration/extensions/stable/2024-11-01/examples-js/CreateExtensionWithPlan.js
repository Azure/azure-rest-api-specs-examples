const { ExtensionsClient } = require("@azure/arm-kubernetesconfiguration-extensions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a new Kubernetes Cluster Extension.
 *
 * @summary create a new Kubernetes Cluster Extension.
 * x-ms-original-file: 2024-11-01/CreateExtensionWithPlan.json
 */
async function createExtensionWithPlan() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId1";
  const client = new ExtensionsClient(credential, subscriptionId);
  const result = await client.extensions.create(
    "rg1",
    "Microsoft.Kubernetes",
    "connectedClusters",
    "clusterName1",
    "azureVote",
    {
      plan: {
        name: "azure-vote-standard",
        product: "azure-vote-standard-offer-id",
        publisher: "Microsoft",
      },
      autoUpgradeMinorVersion: true,
      extensionType: "azure-vote",
      releaseTrain: "Preview",
    },
  );
  console.log(result);
}
