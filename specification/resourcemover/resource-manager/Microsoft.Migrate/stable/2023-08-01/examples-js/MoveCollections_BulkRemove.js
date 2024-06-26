const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Removes the set of move resources included in the request body from move collection. The orchestration is done by service. To aid the user to prerequisite the operation the client can call operation with validateOnly property set to true.
 *
 * @summary Removes the set of move resources included in the request body from move collection. The orchestration is done by service. To aid the user to prerequisite the operation the client can call operation with validateOnly property set to true.
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_BulkRemove.json
 */
async function moveCollectionsBulkRemove() {
  const subscriptionId = process.env["RESOURCEMOVER_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["RESOURCEMOVER_RESOURCE_GROUP"] || "rg1";
  const moveCollectionName = "movecollection1";
  const body = {
    moveResources: [
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1",
    ],
    validateOnly: false,
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const result = await client.moveCollections.beginBulkRemoveAndWait(
    resourceGroupName,
    moveCollectionName,
    options
  );
  console.log(result);
}
