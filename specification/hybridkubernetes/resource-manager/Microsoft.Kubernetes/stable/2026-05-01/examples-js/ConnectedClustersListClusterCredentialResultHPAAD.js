const { ConnectedKubernetesClient } = require("@azure/arm-hybridkubernetes");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets cluster user credentials of the connected cluster with a specified resource group and name.
 *
 * @summary gets cluster user credentials of the connected cluster with a specified resource group and name.
 * x-ms-original-file: 2026-05-01/ConnectedClustersListClusterCredentialResultHPAAD.json
 */
async function listClusterUserCredentialCSPExample() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "1bfbb5d0-917e-4346-9026-1d3b344417f5";
  const client = new ConnectedKubernetesClient(credential, subscriptionId);
  const result = await client.connectedClusterOperations.listClusterUserCredential(
    "k8sc-rg",
    "testCluster",
    { authenticationMethod: "AAD", clientProxy: false },
  );
  console.log(result);
}
