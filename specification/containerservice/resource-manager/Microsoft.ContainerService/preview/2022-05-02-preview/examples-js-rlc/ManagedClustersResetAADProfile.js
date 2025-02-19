const ContainerServiceManagementClient = require("@azure-rest/arm-containerservice").default,
  { getLongRunningPoller } = require("@azure-rest/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Reset the AAD Profile of a managed cluster.
 *
 * @summary Reset the AAD Profile of a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-05-02-preview/examples/ManagedClustersResetAADProfile.json
 */
async function resetAadProfile() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const parameters = {
    body: {
      clientAppID: "clientappid",
      serverAppID: "serverappid",
      serverAppSecret: "serverappsecret",
      tenantID: "tenantid",
    },
  };
  const credential = new DefaultAzureCredential();

  const client = ContainerServiceManagementClient(credential);
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/resetAADProfile",
      subscriptionId,
      resourceGroupName,
      resourceName,
    )
    .post(parameters);
  const poller = await getLongRunningPoller(client, initialResponse);
  const result = poller.pollUntilDone();
  console.log(result);
}

resetAadProfile().catch(console.error);
