const { AgriFoodMgmtClient } = require("@azure/arm-agrifood");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks the name availability of the resource with requested resource name.
 *
 * @summary Checks the name availability of the resource with requested resource name.
 * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Locations_CheckNameAvailability_Available.json
 */
async function locationsCheckNameAvailabilityAvailable() {
  const subscriptionId = "11111111-2222-3333-4444-555555555555";
  const body = {
    name: "newaccountname",
    type: "Microsoft.AgFoodPlatform/farmBeats",
  };
  const credential = new DefaultAzureCredential();
  const client = new AgriFoodMgmtClient(credential, subscriptionId);
  const result = await client.locations.checkNameAvailability(body);
  console.log(result);
}

locationsCheckNameAvailabilityAvailable().catch(console.error);
