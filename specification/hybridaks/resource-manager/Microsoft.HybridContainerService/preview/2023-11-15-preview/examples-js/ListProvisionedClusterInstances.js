const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Hybrid AKS provisioned cluster instances associated with the connected cluster
 *
 * @summary Gets the Hybrid AKS provisioned cluster instances associated with the connected cluster
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2023-11-15-preview/examples/ListProvisionedClusterInstances.json
 */
async function listProvisionedClusterInstances() {
  const connectedClusterResourceUri =
    "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster";
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential);
  const resArray = new Array();
  for await (let item of client.provisionedClusterInstances.list(connectedClusterResourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}
