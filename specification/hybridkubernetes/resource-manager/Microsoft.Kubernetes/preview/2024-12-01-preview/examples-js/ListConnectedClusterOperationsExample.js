const { ConnectedKubernetesClient } = require("@azure/arm-hybridkubernetes");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all of the available API operations for Connected Cluster resource.
 *
 * @summary Lists all of the available API operations for Connected Cluster resource.
 * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/ListConnectedClusterOperationsExample.json
 */
async function listConnectedClusterOperationsExample() {
  const credential = new DefaultAzureCredential();
  const client = new ConnectedKubernetesClient(credential);
  const resArray = new Array();
  for await (const item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
