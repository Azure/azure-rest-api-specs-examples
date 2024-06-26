const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a specified contact profile resource.
 *
 * @summary Deletes a specified contact profile resource.
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/ContactProfileDelete.json
 */
async function deleteContactProfile() {
  const subscriptionId =
    process.env["ORBITAL_SUBSCRIPTION_ID"] || "c1be1141-a7c9-4aac-9608-3c2e2f1152c3";
  const resourceGroupName = process.env["ORBITAL_RESOURCE_GROUP"] || "contoso-Rgp";
  const contactProfileName = "CONTOSO-CP";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const result = await client.contactProfiles.beginDeleteAndWait(
    resourceGroupName,
    contactProfileName
  );
  console.log(result);
}
