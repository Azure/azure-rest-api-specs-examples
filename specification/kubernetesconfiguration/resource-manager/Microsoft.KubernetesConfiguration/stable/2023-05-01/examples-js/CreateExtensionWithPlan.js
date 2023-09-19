const { SourceControlConfigurationClient } = require("@azure/arm-kubernetesconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Kubernetes Cluster Extension.
 *
 * @summary Create a new Kubernetes Cluster Extension.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/CreateExtensionWithPlan.json
 */
async function createExtensionWithPlan() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subId1";
  const resourceGroupName = process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const extensionName = "azureVote";
  const extension = {
    autoUpgradeMinorVersion: true,
    extensionType: "azure-vote",
    plan: {
      name: "azure-vote-standard",
      product: "azure-vote-standard-offer-id",
      publisher: "Microsoft",
    },
    releaseTrain: "Preview",
  };
  const credential = new DefaultAzureCredential();
  const client = new SourceControlConfigurationClient(credential, subscriptionId);
  const result = await client.extensions.beginCreateAndWait(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName,
    extensionName,
    extension
  );
  console.log(result);
}
