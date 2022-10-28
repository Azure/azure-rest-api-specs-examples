const ContainerServiceManagementClient = require("@azure-rest/arm-containerservice").default,
  { getLongRunningPoller } = require("@azure-rest/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to See [Certificate rotation](https://docs.microsoft.com/azure/aks/certificate-rotation) for more details about rotating managed cluster certificates.
 *
 * @summary See [Certificate rotation](https://docs.microsoft.com/azure/aks/certificate-rotation) for more details about rotating managed cluster certificates.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-05-02-preview/examples/ManagedClustersRotateClusterCertificates.json
 */
async function rotateClusterCertificates() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = ContainerServiceManagementClient(credential);
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/rotateClusterCertificates",
      subscriptionId,
      resourceGroupName,
      resourceName
    )
    .post();
  const poller = getLongRunningPoller(client, initialResponse);
  const result = poller.pollUntilDone();
  console.log(result);
}

rotateClusterCertificates().catch(console.error);
