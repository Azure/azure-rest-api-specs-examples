const { LargeInstanceManagementClient } = require("@azure/arm-largeinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches the Tags field of a Azure Large Storage Instance for the specified
subscription, resource group, and instance name.
 *
 * @summary Patches the Tags field of a Azure Large Storage Instance for the specified
subscription, resource group, and instance name.
 * x-ms-original-file: specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/examples/AzureLargeStorageInstance_PatchTags.json
 */
async function azureLargeStorageInstanceUpdateTag() {
  const subscriptionId =
    process.env["LARGEINSTANCE_SUBSCRIPTION_ID"] || "f0f4887f-d13c-4943-a8ba-d7da28d2a3fd";
  const resourceGroupName = process.env["LARGEINSTANCE_RESOURCE_GROUP"] || "myResourceGroup";
  const azureLargeStorageInstanceName = "myALSInstance";
  const tagsParameter = {
    tags: { testkey: "testvalue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new LargeInstanceManagementClient(credential, subscriptionId);
  const result = await client.azureLargeStorageInstanceOperations.update(
    resourceGroupName,
    azureLargeStorageInstanceName,
    tagsParameter,
  );
  console.log(result);
}
