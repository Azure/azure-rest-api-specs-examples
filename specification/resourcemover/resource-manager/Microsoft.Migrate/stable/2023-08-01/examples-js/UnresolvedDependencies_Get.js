const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of unresolved dependencies.
 *
 * @summary Gets a list of unresolved dependencies.
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/UnresolvedDependencies_Get.json
 */
async function unresolvedDependenciesGet() {
  const subscriptionId = process.env["RESOURCEMOVER_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["RESOURCEMOVER_RESOURCE_GROUP"] || "rg1";
  const moveCollectionName = "movecollection1";
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.unresolvedDependencies.list(
    resourceGroupName,
    moveCollectionName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
