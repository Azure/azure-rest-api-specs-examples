const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the specified spacecraft tags.
 *
 * @summary Updates the specified spacecraft tags.
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/SpacecraftUpdateTags.json
 */
async function updateSpacecraftTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const spacecraftName = "AQUA";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const result = await client.spacecrafts.beginUpdateTagsAndWait(
    resourceGroupName,
    spacecraftName,
    parameters
  );
  console.log(result);
}

updateSpacecraftTags().catch(console.error);
