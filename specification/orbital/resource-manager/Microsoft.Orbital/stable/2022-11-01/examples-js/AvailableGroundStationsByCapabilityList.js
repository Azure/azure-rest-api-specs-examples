const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of available ground stations.
 *
 * @summary Returns list of available ground stations.
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/AvailableGroundStationsByCapabilityList.json
 */
async function listOfGroundStationsByCapability() {
  const subscriptionId =
    process.env["ORBITAL_SUBSCRIPTION_ID"] || "c1be1141-a7c9-4aac-9608-3c2e2f1152c3";
  const capability = "EarthObservation";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availableGroundStations.listByCapability(capability)) {
    resArray.push(item);
  }
  console.log(resArray);
}
