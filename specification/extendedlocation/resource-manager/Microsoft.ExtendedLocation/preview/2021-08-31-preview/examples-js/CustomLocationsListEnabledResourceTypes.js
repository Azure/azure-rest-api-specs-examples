const { CustomLocationsManagementClient } = require("@azure/arm-extendedlocation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of the Enabled Resource Types.
 *
 * @summary Gets the list of the Enabled Resource Types.
 * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/CustomLocationsListEnabledResourceTypes.json
 */
async function getCustomLocation() {
  const subscriptionId =
    process.env["EXTENDEDLOCATION_SUBSCRIPTION_ID"] || "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = process.env["EXTENDEDLOCATION_RESOURCE_GROUP"] || "testresourcegroup";
  const resourceName = "customLocation01";
  const credential = new DefaultAzureCredential();
  const client = new CustomLocationsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.customLocations.listEnabledResourceTypes(
    resourceGroupName,
    resourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
