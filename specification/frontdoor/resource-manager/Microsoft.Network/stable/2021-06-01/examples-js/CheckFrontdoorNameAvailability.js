const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the availability of a Front Door resource name.
 *
 * @summary Check the availability of a Front Door resource name.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/CheckFrontdoorNameAvailability.json
 */
async function checkNameAvailability() {
  const checkFrontDoorNameAvailabilityInput = {
    name: "sampleName",
    type: "Microsoft.Network/frontDoors",
  };
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential);
  const result = await client.frontDoorNameAvailability.check(checkFrontDoorNameAvailabilityInput);
  console.log(result);
}
