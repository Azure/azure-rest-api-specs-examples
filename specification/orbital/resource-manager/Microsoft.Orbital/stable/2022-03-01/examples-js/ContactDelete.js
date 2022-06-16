const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a specified contact
 *
 * @summary Deletes a specified contact
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactDelete.json
 */
async function deleteContact() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const spacecraftName = "AQUA";
  const contactName = "contact1";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const result = await client.contacts.beginDeleteAndWait(
    resourceGroupName,
    spacecraftName,
    contactName
  );
  console.log(result);
}

deleteContact().catch(console.error);
