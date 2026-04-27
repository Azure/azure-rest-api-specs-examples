const { ExtensionsClient } = require("@azure/arm-kubernetesconfiguration-extensions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a new Kubernetes Cluster Extension.
 *
 * @summary create a new Kubernetes Cluster Extension.
 * x-ms-original-file: 2025-03-01/CreateExtensionWithPlan.json
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
      extensionType: "azure-vote",
      autoUpgradeMode: "compatible",
      autoUpgradeMinorVersion: true,
      releaseTrain: "Preview",
      plan: {
        name: "azure-vote-standard",
        publisher: "Microsoft",
        product: "azure-vote-standard-offer-id",
      },
    },
  );
  console.log(result);
}
