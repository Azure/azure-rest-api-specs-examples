const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the availability of name for resource
 *
 * @summary Check the availability of name for resource
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/CheckNameAvailability.json
 */
async function nameAvailability() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const nameAvailabilityRequest = {
    name: "name1",
    type: "Microsoft.DevCenter/devcenters",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.checkNameAvailability.execute(nameAvailabilityRequest);
  console.log(result);
}
