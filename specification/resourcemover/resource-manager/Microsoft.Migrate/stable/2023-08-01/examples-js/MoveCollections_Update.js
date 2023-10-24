const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a move collection.
 *
 * @summary Updates a move collection.
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Update.json
 */
async function moveCollectionsUpdate() {
  const subscriptionId = process.env["RESOURCEMOVER_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["RESOURCEMOVER_RESOURCE_GROUP"] || "rg1";
  const moveCollectionName = "movecollection1";
  const body = {
    identity: { type: "SystemAssigned" },
    tags: { key1: "mc1" },
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const result = await client.moveCollections.update(
    resourceGroupName,
    moveCollectionName,
    options
  );
  console.log(result);
}
