const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Move Resource from the move collection.
 *
 * @summary Deletes a Move Resource from the move collection.
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2021-08-01/examples/MoveResources_Delete.json
 */
async function moveResourcesDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const moveCollectionName = "movecollection1";
  const moveResourceName = "moveresourcename1";
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const result = await client.moveResources.beginDeleteAndWait(
    resourceGroupName,
    moveCollectionName,
    moveResourceName
  );
  console.log(result);
}

moveResourcesDelete().catch(console.error);
