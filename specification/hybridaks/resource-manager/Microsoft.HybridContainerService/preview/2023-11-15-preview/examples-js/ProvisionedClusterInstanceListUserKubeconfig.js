const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the AAD user credentials of a provisioned cluster instance used only in direct mode.
 *
 * @summary Lists the AAD user credentials of a provisioned cluster instance used only in direct mode.
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2023-11-15-preview/examples/ProvisionedClusterInstanceListUserKubeconfig.json
 */
async function listClusterUserCredentials() {
  const connectedClusterResourceUri =
    "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster";
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential);
  const result = await client.provisionedClusterInstances.beginListUserKubeconfigAndWait(
    connectedClusterResourceUri
  );
  console.log(result);
}
