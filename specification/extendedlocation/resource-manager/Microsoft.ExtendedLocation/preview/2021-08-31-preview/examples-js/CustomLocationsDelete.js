const { CustomLocationsManagementClient } = require("@azure/arm-extendedlocation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the Custom Location with the specified Resource Name, Resource Group, and Subscription Id.
 *
 * @summary Deletes the Custom Location with the specified Resource Name, Resource Group, and Subscription Id.
 * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/CustomLocationsDelete.json
 */
async function deleteCustomLocation() {
  const subscriptionId =
    process.env["EXTENDEDLOCATION_SUBSCRIPTION_ID"] || "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = process.env["EXTENDEDLOCATION_RESOURCE_GROUP"] || "testresourcegroup";
  const resourceName = "customLocation01";
  const credential = new DefaultAzureCredential();
  const client = new CustomLocationsManagementClient(credential, subscriptionId);
  const result = await client.customLocations.beginDeleteAndWait(resourceGroupName, resourceName);
  console.log(result);
}
