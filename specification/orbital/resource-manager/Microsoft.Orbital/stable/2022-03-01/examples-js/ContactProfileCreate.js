const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a contact profile
 *
 * @summary Creates or updates a contact profile
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfileCreate.json
 */
async function createAContactProfile() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const contactProfileName = "AQUA_DIRECTPLAYBACK_WITH_UPLINK";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const result = await client.contactProfiles.beginCreateOrUpdateAndWait(
    resourceGroupName,
    contactProfileName,
    location
  );
  console.log(result);
}

createAContactProfile().catch(console.error);
