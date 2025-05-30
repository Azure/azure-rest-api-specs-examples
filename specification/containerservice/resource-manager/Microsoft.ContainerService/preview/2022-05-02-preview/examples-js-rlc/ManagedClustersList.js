const ContainerServiceManagementClient = require("@azure-rest/arm-containerservice").default,
  { paginate } = require("@azure-rest/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of managed clusters in the specified subscription.
 *
 * @summary Gets a list of managed clusters in the specified subscription.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-05-02-preview/examples/ManagedClustersList.json
 */
async function listManagedClusters() {
  const subscriptionId = "subid1";
  const credential = new DefaultAzureCredential();
  const client = ContainerServiceManagementClient(credential);
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/managedClusters",
      subscriptionId,
    )
    .get();
  const result = paginate(client, initialResponse);
  const resArray = new Array();
  for await (let item of result) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagedClusters().catch(console.error);
