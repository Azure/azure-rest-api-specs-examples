const { ConnectedKubernetesClient } = require("@azure/arm-hybridkubernetes");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to API to register a new Kubernetes cluster and create or replace a connected cluster tracked resource in Azure Resource Manager (ARM).
 *
 * @summary API to register a new Kubernetes cluster and create or replace a connected cluster tracked resource in Azure Resource Manager (ARM).
 * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/CreateClusterAgentless_KindAWSExample.json
 */
async function createClusterAgentlessKindAwsExample() {
  const subscriptionId =
    process.env["HYBRIDKUBERNETES_SUBSCRIPTION_ID"] || "1bfbb5d0-917e-4346-9026-1d3b344417f5";
  const resourceGroupName = process.env["HYBRIDKUBERNETES_RESOURCE_GROUP"] || "k8sc-rg";
  const clusterName = "testCluster";
  const connectedCluster = {
    identity: { type: "None" },
    kind: "AWS",
    location: "East US",
    properties: {
      agentPublicKeyCertificate: "",
      distribution: "eks",
      infrastructure: "aws",
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new ConnectedKubernetesClient(credential, subscriptionId);
  const result = await client.connectedClusterOperations.beginCreateOrReplaceAndWait(
    resourceGroupName,
    clusterName,
    connectedCluster,
  );
  console.log(result);
}
