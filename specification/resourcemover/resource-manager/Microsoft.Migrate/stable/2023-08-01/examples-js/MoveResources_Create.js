const { ResourceMoverServiceAPI } = require("@azure/arm-resourcemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Move Resource in the move collection.
 *
 * @summary Creates or updates a Move Resource in the move collection.
 * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_Create.json
 */
async function moveResourcesCreate() {
  const subscriptionId = process.env["RESOURCEMOVER_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["RESOURCEMOVER_RESOURCE_GROUP"] || "rg1";
  const moveCollectionName = "movecollection1";
  const moveResourceName = "moveresourcename1";
  const body = {
    properties: {
      dependsOnOverrides: [
        {
          id: "/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/eastusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140",
          targetId:
            "/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/westusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140",
        },
      ],
      resourceSettings: {
        resourceType: "Microsoft.Compute/virtualMachines",
        targetAvailabilitySetId:
          "/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.Compute/availabilitySets/avset1",
        targetAvailabilityZone: "2",
        targetResourceName: "westusvm1",
        targetVmSize: undefined,
        userManagedIdentities: [
          "/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi1",
        ],
      },
      sourceId:
        "/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.Compute/virtualMachines/eastusvm1",
    },
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new ResourceMoverServiceAPI(credential, subscriptionId);
  const result = await client.moveResources.beginCreateAndWait(
    resourceGroupName,
    moveCollectionName,
    moveResourceName,
    options
  );
  console.log(result);
}
