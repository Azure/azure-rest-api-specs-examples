const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a specified spacecraft resource.
 *
 * @summary Deletes a specified spacecraft resource.
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/SpacecraftDelete.json
 */
async function deleteSpacecraft() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const spacecraftName = "AQUA";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const result = await client.spacecrafts.beginDeleteAndWait(resourceGroupName, spacecraftName);
  console.log(result);
}

deleteSpacecraft().catch(console.error);
