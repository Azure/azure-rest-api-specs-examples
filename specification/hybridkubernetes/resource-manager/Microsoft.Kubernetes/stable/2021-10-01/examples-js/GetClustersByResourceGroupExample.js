const { ConnectedKubernetesClient } = require("@azure/arm-hybridkubernetes");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to API to enumerate registered connected K8s clusters under a Resource Group
 *
 * @summary API to enumerate registered connected K8s clusters under a Resource Group
 * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/stable/2021-10-01/examples/GetClustersByResourceGroupExample.json
 */
async function getClustersExample() {
  const subscriptionId = "1bfbb5d0-917e-4346-9026-1d3b344417f5";
  const resourceGroupName = "k8sc-rg";
  const credential = new DefaultAzureCredential();
  const client = new ConnectedKubernetesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.connectedClusterOperations.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getClustersExample().catch(console.error);
