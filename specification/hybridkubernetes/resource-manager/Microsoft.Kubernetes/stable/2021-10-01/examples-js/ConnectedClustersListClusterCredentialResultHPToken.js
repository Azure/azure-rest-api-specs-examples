const { ConnectedKubernetesClient } = require("@azure/arm-hybridkubernetes");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets cluster user credentials of the connected cluster with a specified resource group and name.
 *
 * @summary Gets cluster user credentials of the connected cluster with a specified resource group and name.
 * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/stable/2021-10-01/examples/ConnectedClustersListClusterCredentialResultHPToken.json
 */
async function listClusterUserCredentialNonAadCspExample() {
  const subscriptionId = "1bfbb5d0-917e-4346-9026-1d3b344417f5";
  const resourceGroupName = "k8sc-rg";
  const clusterName = "testCluster";
  const properties = {
    authenticationMethod: "Token",
    clientProxy: false,
  };
  const credential = new DefaultAzureCredential();
  const client = new ConnectedKubernetesClient(credential, subscriptionId);
  const result = await client.connectedClusterOperations.listClusterUserCredential(
    resourceGroupName,
    clusterName,
    properties
  );
  console.log(result);
}

listClusterUserCredentialNonAadCspExample().catch(console.error);
