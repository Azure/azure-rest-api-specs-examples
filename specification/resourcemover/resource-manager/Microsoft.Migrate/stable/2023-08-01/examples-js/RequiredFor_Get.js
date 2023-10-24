const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of the move resources for which an arm resource is required for.
 *
 * @summary List of the move resources for which an arm resource is required for.
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/RequiredFor_Get.json
 */
async function requiredForGet() {
  const subscriptionId = process.env["RESOURCEMOVER_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["RESOURCEMOVER_RESOURCE_GROUP"] || "rg1";
  const moveCollectionName = "movecollection1";
  const sourceId =
    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/nic1";
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const result = await client.moveCollections.listRequiredFor(
    resourceGroupName,
    moveCollectionName,
    sourceId
  );
  console.log(result);
}
