const { ConnectedKubernetesClient } = require("@azure/arm-hybridkubernetes");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to API to register a new Kubernetes cluster and create or replace a connected cluster tracked resource in Azure Resource Manager (ARM).
 *
 * @summary API to register a new Kubernetes cluster and create or replace a connected cluster tracked resource in Azure Resource Manager (ARM).
 * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/CreateCluster_KindExample.json
 */
async function createClusterKindExample() {
  const subscriptionId =
    process.env["HYBRIDKUBERNETES_SUBSCRIPTION_ID"] || "1bfbb5d0-917e-4346-9026-1d3b344417f5";
  const resourceGroupName = process.env["HYBRIDKUBERNETES_RESOURCE_GROUP"] || "k8sc-rg";
  const clusterName = "testCluster";
  const connectedCluster = {
    identity: { type: "SystemAssigned" },
    kind: "ProvisionedCluster",
    location: "East US",
    properties: {
      aadProfile: {
        adminGroupObjectIDs: ["56f988bf-86f1-41af-91ab-2d7cd011db47"],
        enableAzureRbac: true,
        tenantID: "82f988bf-86f1-41af-91ab-2d7cd011db47",
      },
      agentPublicKeyCertificate: "",
      arcAgentProfile: {
        agentAutoUpgrade: "Enabled",
        desiredAgentVersion: "0.1.0",
        systemComponents: [{ type: "Strato", majorVersion: 0, userSpecifiedVersion: "0.1.1" }],
      },
      azureHybridBenefit: "NotApplicable",
      distribution: "AKS",
      distributionVersion: "1.0",
      oidcIssuerProfile: { enabled: true },
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
