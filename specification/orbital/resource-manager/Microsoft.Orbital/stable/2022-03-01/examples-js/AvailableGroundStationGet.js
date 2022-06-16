const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified  available ground station
 *
 * @summary Gets the specified  available ground station
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/AvailableGroundStationGet.json
 */
async function getGroundStation() {
  const subscriptionId = "subid";
  const groundStationName = "westus_gs1";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const result = await client.availableGroundStations.get(groundStationName);
  console.log(result);
}

getGroundStation().catch(console.error);
