const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Move Resource.
 *
 * @summary Gets the Move Resource.
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_Get.json
 */
async function moveResourcesGet() {
  const subscriptionId = process.env["RESOURCEMOVER_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["RESOURCEMOVER_RESOURCE_GROUP"] || "rg1";
  const moveCollectionName = "movecollection1";
  const moveResourceName = "moveresourcename1";
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const result = await client.moveResources.get(
    resourceGroupName,
    moveCollectionName,
    moveResourceName
  );
  console.log(result);
}
