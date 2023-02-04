const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified contact Profile in a specified resource group
 *
 * @summary Gets the specified contact Profile in a specified resource group
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfileGet.json
 */
async function getAContactProfile() {
  const subscriptionId = process.env["ORBITAL_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["ORBITAL_RESOURCE_GROUP"] || "rg1";
  const contactProfileName = "AQUA_DIRECTPLAYBACK_WITH_UPLINK";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const result = await client.contactProfiles.get(resourceGroupName, contactProfileName);
  console.log(result);
}
