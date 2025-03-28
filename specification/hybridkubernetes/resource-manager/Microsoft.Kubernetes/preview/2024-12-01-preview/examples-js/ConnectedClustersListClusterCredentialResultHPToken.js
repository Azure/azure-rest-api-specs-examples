const { ConnectedKubernetesClient } = require("@azure/arm-hybridkubernetes");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets cluster user credentials of the connected cluster with a specified resource group and name.
 *
 * @summary Gets cluster user credentials of the connected cluster with a specified resource group and name.
 * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/ConnectedClustersListClusterCredentialResultHPToken.json
 */
async function listClusterUserCredentialNonAadCspExample() {
  const subscriptionId =
    process.env["HYBRIDKUBERNETES_SUBSCRIPTION_ID"] || "1bfbb5d0-917e-4346-9026-1d3b344417f5";
  const resourceGroupName = process.env["HYBRIDKUBERNETES_RESOURCE_GROUP"] || "k8sc-rg";
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
    properties,
  );
  console.log(result);
}
