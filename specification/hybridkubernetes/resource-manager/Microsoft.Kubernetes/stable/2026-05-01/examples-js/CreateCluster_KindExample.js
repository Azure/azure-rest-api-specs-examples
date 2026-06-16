const { ConnectedKubernetesClient } = require("@azure/arm-hybridkubernetes");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to aPI to register a new Kubernetes cluster and create or replace a connected cluster tracked resource in Azure Resource Manager (ARM).
 *
 * @summary aPI to register a new Kubernetes cluster and create or replace a connected cluster tracked resource in Azure Resource Manager (ARM).
 * x-ms-original-file: 2026-05-01/CreateCluster_KindExample.json
 */
async function createClusterKindExample() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "1bfbb5d0-917e-4346-9026-1d3b344417f5";
  const client = new ConnectedKubernetesClient(credential, subscriptionId);
  const result = await client.connectedClusterOperations.createOrReplace("k8sc-rg", "testCluster", {
    identity: { type: "SystemAssigned" },
    kind: "ProvisionedCluster",
    location: "East US",
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
    tags: {},
  });
  console.log(result);
}
