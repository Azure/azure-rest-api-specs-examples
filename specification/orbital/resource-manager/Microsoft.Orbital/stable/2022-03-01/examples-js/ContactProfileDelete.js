const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a specified contact profile resource.
 *
 * @summary Deletes a specified contact profile resource.
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfileDelete.json
 */
async function deleteContactProfile() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const contactProfileName = "AQUA_DIRECTPLAYBACK_WITH_UPLINK";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const result = await client.contactProfiles.beginDeleteAndWait(
    resourceGroupName,
    contactProfileName
  );
  console.log(result);
}

deleteContactProfile().catch(console.error);
