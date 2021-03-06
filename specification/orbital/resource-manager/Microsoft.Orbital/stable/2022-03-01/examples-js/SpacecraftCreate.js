const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a spacecraft resource
 *
 * @summary Creates or updates a spacecraft resource
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/SpacecraftCreate.json
 */
async function createASpacecraft() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const spacecraftName = "AQUA";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const result = await client.spacecrafts.beginCreateOrUpdateAndWait(
    resourceGroupName,
    spacecraftName,
    location
  );
  console.log(result);
}

createASpacecraft().catch(console.error);
