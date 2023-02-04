const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified spacecraft in a specified resource group
 *
 * @summary Gets the specified spacecraft in a specified resource group
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/SpacecraftGet.json
 */
async function getSpacecraft() {
  const subscriptionId = process.env["ORBITAL_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["ORBITAL_RESOURCE_GROUP"] || "rg1";
  const spacecraftName = "AQUA";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const result = await client.spacecrafts.get(resourceGroupName, spacecraftName);
  console.log(result);
}
