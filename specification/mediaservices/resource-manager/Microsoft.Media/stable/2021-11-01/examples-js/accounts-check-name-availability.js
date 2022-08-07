const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks whether the Media Service resource name is available.
 *
 * @summary Checks whether the Media Service resource name is available.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/accounts-check-name-availability.json
 */
async function checkNameAvailability() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const locationName = "japanwest";
  const parameters = {
    name: "contosotv",
    type: "videoAnalyzers",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.locations.checkNameAvailability(locationName, parameters);
  console.log(result);
}

checkNameAvailability().catch(console.error);
