```javascript
const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of available ground stations
 *
 * @summary Returns list of available ground stations
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/AvailableGroundStationsByCapabilityList.json
 */
async function listOfGroundStationsByCapability() {
  const subscriptionId = "subId";
  const capability = "EarthObservation";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availableGroundStations.listByCapability(capability)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listOfGroundStationsByCapability().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-orbital_1.0.0/sdk/orbital/arm-orbital/README.md) on how to add the SDK to your project and authenticate.
