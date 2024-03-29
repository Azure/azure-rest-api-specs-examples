const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the specified spacecraft tags.
 *
 * @summary Updates the specified spacecraft tags.
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/SpacecraftUpdateTags.json
 */
async function updateSpacecraftTags() {
  const subscriptionId =
    process.env["ORBITAL_SUBSCRIPTION_ID"] || "c1be1141-a7c9-4aac-9608-3c2e2f1152c3";
  const resourceGroupName = process.env["ORBITAL_RESOURCE_GROUP"] || "contoso-Rgp";
  const spacecraftName = "CONTOSO_SAT";
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
