const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Rotates the service account signing keys of a managed cluster.
 *
 * @summary Rotates the service account signing keys of a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-07-02-preview/examples/ManagedClustersRotateServiceAccountSigningKeys.json
 */
async function rotateClusterServiceAccountSigningKeys() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.beginRotateServiceAccountSigningKeysAndWait(
    resourceGroupName,
    resourceName
  );
  console.log(result);
}

rotateClusterServiceAccountSigningKeys().catch(console.error);
