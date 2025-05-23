const { DependencyMapClient } = require("@azure/arm-dependencymap");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a MapsResource
 *
 * @summary create a MapsResource
 * x-ms-original-file: 2025-01-31-preview/Maps_CreateOrUpdate.json
 */
async function mapsCreateOrUpdateGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "D6E58BDB-45F1-41EC-A884-1FC945058848";
  const client = new DependencyMapClient(credential, subscriptionId);
  const result = await client.maps.createOrUpdate("rgdependencyMap", "mapsTest1", {
    properties: {},
    tags: {},
    location: "wjtzelgfcmswfeflfltwxqveo",
  });
  console.log(result);
}
