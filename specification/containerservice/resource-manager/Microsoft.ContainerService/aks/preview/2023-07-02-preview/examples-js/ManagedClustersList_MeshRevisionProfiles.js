const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Contains extra metadata on each revision, including supported revisions, cluster compatibility and available upgrades
 *
 * @summary Contains extra metadata on each revision, including supported revisions, cluster compatibility and available upgrades
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-07-02-preview/examples/ManagedClustersList_MeshRevisionProfiles.json
 */
async function listMeshRevisionProfilesInALocation() {
  const subscriptionId = process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "subid1";
  const location = "location1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedClusters.listMeshRevisionProfiles(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
