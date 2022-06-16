const { ExternalIdentitiesConfigurationClient } = require("@azure/arm-azureadexternalidentities");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks the availability and validity of a domain name for the tenant.
 *
 * @summary Checks the availability and validity of a domain name for the tenant.
 * x-ms-original-file: specification/cpim/resource-manager/Microsoft.AzureActiveDirectory/stable/2021-04-01/examples/checkNameAvailability-available.json
 */
async function checkNameAvailabilityAvailable() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const checkNameAvailabilityRequestBody = {
    name: "constoso.onmicrosoft.com",
    countryCode: "US",
  };
  const options = {
    checkNameAvailabilityRequestBody,
  };
  const credential = new DefaultAzureCredential();
  const client = new ExternalIdentitiesConfigurationClient(credential, subscriptionId);
  const result = await client.b2CTenants.checkNameAvailability(options);
  console.log(result);
}

checkNameAvailabilityAvailable().catch(console.error);
