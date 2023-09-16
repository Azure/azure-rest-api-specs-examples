const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Contains extra metadata on the revision, including supported revisions, cluster compatibility and available upgrades
 *
 * @summary Contains extra metadata on the revision, including supported revisions, cluster compatibility and available upgrades
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-07-02-preview/examples/ManagedClustersGet_MeshRevisionProfile.json
 */
async function getAMeshRevisionProfileForAMeshMode() {
  const subscriptionId = process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "subid1";
  const location = "location1";
  const mode = "istio";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.getMeshRevisionProfile(location, mode);
  console.log(result);
}
