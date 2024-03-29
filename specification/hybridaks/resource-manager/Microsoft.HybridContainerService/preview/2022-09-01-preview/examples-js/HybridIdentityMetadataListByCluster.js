const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the hybrid identity metadata proxy resource in a cluster.
 *
 * @summary Lists the hybrid identity metadata proxy resource in a cluster.
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/HybridIdentityMetadataListByCluster.json
 */
async function hybridIdentityMetadataListByCluster() {
  const subscriptionId =
    process.env["HYBRIDCONTAINERSERVICE_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["HYBRIDCONTAINERSERVICE_RESOURCE_GROUP"] || "testrg";
  const resourceName = "ContosoTargetCluster";
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.hybridIdentityMetadataOperations.listByCluster(
    resourceGroupName,
    resourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
