const { StorageActionsManagementClient } = require("@azure/arm-storageactions");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update storage task properties
 *
 * @summary Update storage task properties
 * x-ms-original-file: specification/storageactions/resource-manager/Microsoft.StorageActions/stable/2023-01-01/examples/storageTasksCrud/PatchStorageTask.json
 */
async function patchStorageTask() {
  const subscriptionId =
    process.env["STORAGEACTIONS_SUBSCRIPTION_ID"] || "1f31ba14-ce16-4281-b9b4-3e78da6e1616";
  const resourceGroupName = process.env["STORAGEACTIONS_RESOURCE_GROUP"] || "res4228";
  const storageTaskName = "mytask1";
  const parameters = {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/1f31ba14Ce164281B9b43e78da6e1616/resourceGroups/res4228/providers/MicrosoftManagedIdentity/userAssignedIdentities/myUserAssignedIdentity":
          {},
      },
    },
    properties: {
      description: "My Storage task",
      action: {
        else: {
          operations: [{ name: "DeleteBlob", onFailure: "break", onSuccess: "continue" }],
        },
        if: {
          condition: "[[equals(AccessTier, 'Cool')]]",
          operations: [
            {
              name: "SetBlobTier",
              onFailure: "break",
              onSuccess: "continue",
              parameters: { tier: "Hot" },
            },
          ],
        },
      },
      enabled: true,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageActionsManagementClient(credential, subscriptionId);
  const result = await client.storageTasks.beginUpdateAndWait(
    resourceGroupName,
    storageTaskName,
    parameters,
  );
  console.log(result);
}
