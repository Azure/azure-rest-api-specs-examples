const { DeviceUpdate } = require("@azure/arm-deviceupdate");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks ADU resource name availability.
 *
 * @summary Checks ADU resource name availability.
 * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/CheckNameAvailability_AlreadyExists.json
 */
async function checkNameAvailabilityAlreadyExists() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const request = {
    name: "contoso",
    type: "Microsoft.DeviceUpdate/accounts",
  };
  const credential = new DefaultAzureCredential();
  const client = new DeviceUpdate(credential, subscriptionId);
  const result = await client.checkNameAvailability(request);
  console.log(result);
}

checkNameAvailabilityAlreadyExists().catch(console.error);
