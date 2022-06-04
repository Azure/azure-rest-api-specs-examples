Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-orbital_1.0.0/sdk/orbital/arm-orbital/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return list of available contacts
 *
 * @summary Return list of available contacts
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/AvailableContactsList.json
 */
async function listOfContact() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgName";
  const spacecraftName = "AQUA";
  const contactProfile = {
    id: "/subscriptions/subId/resourceGroups/rg/Microsoft.Orbital/contactProfiles/AQUA_DIRECTPLAYBACK_WITH_UPLINK",
  };
  const groundStationName = "westus_gs1";
  const startTime = new Date("2020-07-16T05:40:21.00Z");
  const endTime = new Date("2020-07-17T23:49:40.00Z");
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.spacecrafts.beginListAvailableContactsAndWait(
    resourceGroupName,
    spacecraftName,
    contactProfile,
    groundStationName,
    startTime,
    endTime
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listOfContact().catch(console.error);
```
