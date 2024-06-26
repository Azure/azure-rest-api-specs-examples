const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check Name Availability for local uniqueness
 *
 * @summary Check Name Availability for local uniqueness
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/proxy/CheckNameAvailabilityForLocalUniqueness.json
 */
async function checkLocalNameAvailability() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const location = "eastus2euap";
  const checkNameAvailability = {
    name: "MyAccount",
    type: "Microsoft.MixedReality/spatialAnchorsAccounts",
  };
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const result = await client.checkNameAvailabilityLocal(location, checkNameAvailability);
  console.log(result);
}

checkLocalNameAvailability().catch(console.error);
