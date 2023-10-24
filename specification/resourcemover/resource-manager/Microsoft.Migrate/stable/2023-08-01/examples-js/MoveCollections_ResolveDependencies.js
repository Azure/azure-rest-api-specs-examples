const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Computes, resolves and validate the dependencies of the moveResources in the move collection.
 *
 * @summary Computes, resolves and validate the dependencies of the moveResources in the move collection.
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_ResolveDependencies.json
 */
async function moveCollectionsResolveDependencies() {
  const subscriptionId = process.env["RESOURCEMOVER_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["RESOURCEMOVER_RESOURCE_GROUP"] || "rg1";
  const moveCollectionName = "movecollection1";
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const result = await client.moveCollections.beginResolveDependenciesAndWait(
    resourceGroupName,
    moveCollectionName
  );
  console.log(result);
}
