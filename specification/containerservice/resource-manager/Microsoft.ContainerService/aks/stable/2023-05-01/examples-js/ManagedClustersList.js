const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of managed clusters in the specified subscription.
 *
 * @summary Gets a list of managed clusters in the specified subscription.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersList.json
 */
async function listManagedClusters() {
  const subscriptionId = process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "subid1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedClusters.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
