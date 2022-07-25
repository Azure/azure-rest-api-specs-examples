const { CustomLocationsManagementClient } = require("@azure/arm-extendedlocation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of Custom Locations in the specified subscription. The operation returns properties of each Custom Location
 *
 * @summary Gets a list of Custom Locations in the specified subscription. The operation returns properties of each Custom Location
 * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/CustomLocationsListBySubscription.json
 */
async function listCustomLocationsBySubscription() {
  const subscriptionId = "11111111-2222-3333-4444-555555555555";
  const credential = new DefaultAzureCredential();
  const client = new CustomLocationsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.customLocations.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listCustomLocationsBySubscription().catch(console.error);
